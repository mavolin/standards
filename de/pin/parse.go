package pin

import (
	"errors"
	"math"
	"strings"
)

var (
	ErrLength = errors.New("de/pin: pension insurance numbers must be 12 digits long")

	ErrAreaCodeSyntax    = errors.New("de/pin: area codes must only contain digits")
	ErrAreaCodeInvalid   = errors.New("de/pin: invalid area code")
	ErrBirthDay          = errors.New("de/pin: the birthday must only contain digits")
	ErrBirthMonth        = errors.New("de/pin: the birth month must be a number between '01' and '12'")
	ErrBirthYear         = errors.New("de/pin: the birth year must contain only digits")
	ErrLastNameLetter    = errors.New("de/pin: the last name letter must be an ascii letter")
	ErrSerialNumber      = errors.New("de/pin: the serial number must contain only digits")
	ErrCheckDigitSyntax  = errors.New("de/pin: the check digit must be a digit")
	ErrCheckDigitInvalid = errors.New("de/pin: invalid check digit")
)

// Parse parses the passed pension insurance number.
//
// It ignores spaces and '/'.
//
// Parse is case-insensitive, however, the case of letters in the returned
// PensionInsuranceNumber will always be uppercase.
//
// If Parse returns without an error, the pension insurance number is
// considered syntactically valid.
func Parse(s string) (PensionInsuranceNumber, error) {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ReplaceAll(s, "/", "")

	if len(s) != 12 {
		return PensionInsuranceNumber{}, ErrLength
	}

	var pin PensionInsuranceNumber

	untypAreaCode, ok := parseTwoDigits(s[:2])
	if !ok {
		return PensionInsuranceNumber{}, ErrAreaCodeSyntax
	}

	pin.AreaCode = AreaCode(untypAreaCode)
	if !pin.AreaCode.IsValid() {
		return PensionInsuranceNumber{}, ErrAreaCodeInvalid
	}

	// can be 0-99, see doc of pin.BirthDay
	pin.BirthDay, ok = parseTwoDigits(s[2:4])
	if !ok {
		return PensionInsuranceNumber{}, ErrBirthDay
	}

	pin.BirthMonth, ok = parseTwoDigits(s[4:6])
	if !ok || pin.BirthMonth < 1 || pin.BirthMonth > 12 {
		return PensionInsuranceNumber{}, ErrBirthMonth
	}

	pin.BirthYear, ok = parseTwoDigits(s[6:8])
	if !ok {
		return PensionInsuranceNumber{}, ErrBirthYear
	}

	pin.LastNameLetter = rune(s[8])
	if pin.LastNameLetter >= 'a' && pin.LastNameLetter <= 'z' {
		pin.LastNameLetter -= 'a' - 'A' // turn into uppercase
	} else if pin.LastNameLetter < 'A' || pin.LastNameLetter > 'Z' {
		return PensionInsuranceNumber{}, ErrLastNameLetter
	}

	pin.SerialNumber, ok = parseTwoDigits(s[9:11])
	if !ok {
		return PensionInsuranceNumber{}, ErrSerialNumber
	}

	pin.CheckDigit, ok = parseOneDigit(s[11:])
	if !ok {
		return PensionInsuranceNumber{}, ErrCheckDigitSyntax
	}

	if pin.CheckDigit != calcCheckDigit(pin) {
		return PensionInsuranceNumber{}, ErrCheckDigitInvalid
	}

	return pin, nil
}

// save ourselves pre-checks for '_', etc. when using strconv.ParseUint
func parseOneDigit(s string) (uint8, bool) {
	return parseDigit(s[0])
}

func parseTwoDigits(s string) (uint8, bool) {
	digit1, ok := parseDigit(s[0])
	if !ok {
		return 0, false
	}

	digit2, ok := parseDigit(s[1])
	if !ok {
		return 0, false
	}

	return digit1*10 + digit2, true
}

func parseDigit(digit byte) (uint8, bool) {
	if digit < '0' || digit > '9' {
		return 0, false
	}

	return digit - '0', true
}

func calcCheckDigit(pin PensionInsuranceNumber) uint8 {
	// https://de.wikipedia.org/wiki/Versicherungsnummer#Berechnung_der_Pr%C3%BCfziffer
	// 2023-01-22

	var sum int

	sum += int(digitSum(2 * nthDigit(uint8(pin.AreaCode), 2)))
	sum += int(digitSum(1 * nthDigit(uint8(pin.AreaCode), 1)))
	sum += int(digitSum(2 * nthDigit(pin.BirthDay, 2)))
	sum += int(digitSum(5 * nthDigit(pin.BirthDay, 1)))
	sum += int(digitSum(7 * nthDigit(pin.BirthMonth, 2)))
	sum += int(digitSum(1 * nthDigit(pin.BirthMonth, 1)))
	sum += int(digitSum(2 * nthDigit(pin.BirthYear, 2)))
	sum += int(digitSum(1 * nthDigit(pin.BirthYear, 1)))

	numericLetter := uint8(pin.LastNameLetter - 'A' + 1)
	sum += int(digitSum(2 * nthDigit(numericLetter, 2)))
	sum += int(digitSum(1 * nthDigit(numericLetter, 1)))

	sum += int(digitSum(2 * nthDigit(pin.SerialNumber, 2)))
	sum += int(digitSum(1 * nthDigit(pin.SerialNumber, 1)))

	return uint8(sum % 10)
}

// nthDigit returns the nth digit of n, where n=1 would return the rightmost
// digit.
func nthDigit(num uint8, n int) uint8 {
	return (num / (uint8)(math.Pow10(n-1))) % 10
}

func digitSum(n uint8) uint8 {
	return n/100 + n/10 + n%10
}

func IsValid(s string) bool {
	_, err := Parse(s)
	return err == nil
}
