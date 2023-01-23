package iban

import (
	"strings"

	"github.com/mavolin/standards/iso3166"
)

var checksumFuncs = map[iso3166.Alpha2Code]func(IBAN) bool{
	// https://en.wikipedia.org/wiki/International_Bank_Account_Number#National_check_digits
	// 2023-01-22
	// https://bank-code.net/iban/country-list
	// 2021-01-22
	//
	// Quick note on the Wikipedia source:
	// What they are saying in the paragraph above the table is not quite right
	// (or at the very least incredibly misleading):
	// Wiki says that the default case, i.e. when there is no comment, is that
	// the algorithm is used on the account number.
	// However, that is not true.
	// Although, I only have a single test IBAN for each country, for each that
	// uses ISO 7064 MOD-97-10 with complement 98 - r, I could only get the
	// correct national check digit when I calculated the checksum for the
	// entire BBAN and replaced the checksum with "00", instead of just the
	// iban.AccountNumber.
	// That seems like too big of a coincidence to truly be one.

	iso3166.AL: bankCodeBranchCode_weighted9731Mod10_10MinR,
	iso3166.BE: bankCodeAccountNumber_iso7064Mod9710_98MinR,
	iso3166.BA: bankCodeBranchCodeAccountNumber00_iso7064Mod9710_98MinR,
	iso3166.HR: func(iban IBAN) bool {
		check := iso7064Mod1110(iban.BankCode)
		check = 11 - check
		if check != 9 {
			return false
		}

		check = iso7064Mod1110(iban.AccountNumber)
		check = 11 - check
		if check != 9 {
			return false
		}

		return true
	},
	iso3166.CZ: czech,
	iso3166.TL: bankCodeBranchCodeAccountNumber00_iso7064Mod9710_98MinR,
	iso3166.EE: accountNumber_weighted37137137137Mod10_10MinR,
	iso3166.FI: bankCodeAccountNumber_luhn_10MinR,
	// wikipedia's algorithm doesn't work again, luckily the source they linked
	// uses a different algorithm (i'm just as confused as you are), that one
	// worked
	iso3166.FR: france,
	iso3166.HU: accountNumber_weighted9731Mod10_10MinR,
	iso3166.IS: func(iban IBAN) bool {
		check := weighted(iban.OwnerIdentificationNumber[:8], 11, 3, 2, 7, 6, 5, 4, 3, 2)
		if check == 1 {
			return false
		} else if check != 0 {
			check = 11 - check
		}

		return iban.OwnerIdentificationNumber[8] == byte('0'+check)
	},
	// todo: italy
	iso3166.MK: bankCodeBranchCodeAccountNumber00_iso7064Mod9710_98MinR,
	iso3166.MC: france,
	iso3166.ME: bankCodeBranchCodeAccountNumber00_iso7064Mod9710_98MinR,
	iso3166.NO: func(iban IBAN) bool {
		var check int
		if strings.HasPrefix(iban.AccountNumber, "00") {
			check = weighted(iban.AccountNumber[2:], 11, 5, 4, 3, 2)
		} else {
			check = weighted(iban.BankCode+iban.AccountNumber, 11, 5, 4, 3, 2, 7, 6, 5, 4, 3, 2)
		}

		if check == 1 {
			return false
		} else if check != 0 {
			check = 11 - check
		}

		return iban.NationalChecksum == string(rune('0'+check))
	},
	iso3166.PL: bankCodeBranchCode_weighted3971_10MinR,
	iso3166.PT: bankCodeBranchCodeAccountNumber00_iso7064Mod9710_98MinR,
	// todo: san marino uses the same algorithm as italy
	iso3166.RS: bankCodeBranchCodeAccountNumber00_iso7064Mod9710_98MinR,
	iso3166.SK: czech,
	iso3166.ES: func(iban IBAN) bool {
		check1 := weighted(iban.BankCode+iban.BranchCode, 11, 4, 8, 5, 10, 9, 7, 3, 6)
		if check1 > 1 {
			check1 = 11 - check1
		}

		check2 := weighted(iban.AccountNumber, 11, 1, 2, 4, 8, 5, 10, 9, 7, 3, 6)
		if check2 > 1 {
			check2 = 11 - check2
		}

		return iban.NationalChecksum == string(rune('0'+check1))+string(rune('0'+check2))
	},
	iso3166.SI: bankCodeBranchCodeAccountNumber00_iso7064Mod9710_98MinR,
}

// naming: concatenatedFields_algoName_complement

func bankCodeBranchCodeAccountNumber00_iso7064Mod9710_98MinR(iban IBAN) bool {
	check := iso7064Mod9710(iban.BankCode + iban.BranchCode + iban.AccountNumber + "00")
	check = 98 - check
	return iban.NationalChecksum == uint8ToTwoDigitString(uint8(check))
}

func bankCodeAccountNumber_iso7064Mod9710_98MinR(iban IBAN) bool {
	check := iso7064Mod9710(iban.BankCode + iban.AccountNumber)
	if check == 0 {
		check = 97
	}

	return iban.NationalChecksum == uint8ToTwoDigitString(uint8(check))
}

func bankCodeBranchCode_weighted9731Mod10_10MinR(iban IBAN) bool {
	check := weighted(iban.BankCode+iban.BranchCode, 10, 9, 7, 3, 1)
	if check != 0 {
		check = 10 - check
	}

	return iban.NationalChecksum == string(rune('0'+check))
}

func bankCodeBranchCode_weighted3971_10MinR(iban IBAN) bool {
	check := weighted(iban.BankCode+iban.BranchCode, 10, 3, 9, 7, 1)
	if check != 0 {
		check = 10 - check
	}

	return iban.NationalChecksum == string(rune('0'+check))
}

func accountNumber_weighted9731Mod10_10MinR(iban IBAN) bool {
	check := weighted(iban.AccountNumber, 10, 9, 7, 3, 1)
	if check != 0 {
		check = 10 - check
	}

	return iban.NationalChecksum == string(rune('0'+check))
}

func accountNumber_weighted37137137137Mod10_10MinR(iban IBAN) bool {
	check := weighted(iban.AccountNumber, 10, 3, 7, 1, 3, 7, 1, 3, 7, 1, 3, 7)
	if check != 0 {
		check = 10 - check
	}

	return iban.NationalChecksum == string(rune(check+'0'))
}

func bankCodeAccountNumber_luhn_10MinR(iban IBAN) bool {
	check := luhn(iban.BankCode + iban.AccountNumber)
	if check != 0 {
		check = 10 - check
	}

	return iban.NationalChecksum == string(rune('0'+check))
}

func czech(iban IBAN) bool {
	r := weighted(iban.BranchCode, 11, 10, 5, 8, 4, 2, 1)
	if r != 0 {
		return false
	}

	r = weighted(iban.AccountNumber, 11, 6, 3, 7, 9, 10, 5, 8, 4, 2, 1)
	if r != 0 {
		return false
	}

	return true
}

func france(iban IBAN) bool {
	s := franceTransliterate(iban.BankCode + iban.BranchCode + iban.AccountNumber + iban.NationalChecksum)
	check := iso7064Mod9710(s)
	return check == 0
}

func franceTransliterate(s string) string {
	for i := 0; i < len(s); i++ {
		r := s[i]
		switch {
		case r >= 'A' && r <= 'I':
			s = s[:i] + string(rune(r-'A'+'1')) + s[i+1:]
		case r >= 'J' && r <= 'N':
			s = s[:i] + string(rune(r-'J'+'1')) + s[i+1:]
		case r >= 'S' && r <= 'Z':
			s = s[:i] + string(rune(r-'S'+'2')) + s[i+1:]
		}
	}

	return s
}

// ============================================================================
// Algorithms
// ======================================================================================

func iso7064Mod9710(s string) int {
	var remainder int
	for _, r := range s {
		remainder = remainder*10 + int(r-'0')
		remainder %= 97
	}

	return remainder
}

// implements ISO 7064 MOD-11-10
func iso7064Mod1110(s string) int {
	n := 10
	for _, r := range s {
		n += int(r - '0')
		n %= 10
		if n == 0 {
			n = 10
		}

		n *= 2
		n %= 11
	}

	return n
}

func weighted(s string, mod int, weights ...int) int {
	var sum uint64

	for i, r := range s {
		sum += uint64(int(r-'0') * weights[i%len(weights)])
	}

	return int(sum) % mod
}

func luhn(s string) int {
	var sum uint64

	for i := len(s) - 1; i >= 0; i-- {
		n := int(s[i] - '0')
		if i%2 == 0 {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}

		sum += uint64(n)
	}

	return (10 - int(sum%10)) % 10
}
