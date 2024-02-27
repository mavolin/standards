// Package iban provides parsing and validation for International Bank Account
// Numbers (IBAN).
//
// It can validate the country code, checksum, and country-specific length and
// checksum (if the country has one).
//
// For countries for which no information is available, it only validates the
// country code and checksum.
package iban

import (
	"encoding"
	"strings"

	"github.com/mavolin/standards/iso3166"
)

const maxLen = 34

// IBAN represents an International Bank Account Number.
type IBAN struct {
	// CountryCode is the ISO 3166-1 alpha-2 country code of the IBAN.
	CountryCode iso3166.Alpha2Code
	// Checksum is the checksum of the IBAN.
	Checksum uint8
	// BBAN is the Basic Bank Account Number of the IBAN.
	BBAN string

	// Country-specific fields; might not always be present.
	// Filled according to
	// https://en.wikipedia.org/wiki/International_Bank_Account_Number#IBAN_formats_by_country

	BankCode                  string
	BranchCode                string
	AccountType               string
	OwnerAccountNumber        string
	OwnerIdentificationNumber string // e.g. kennitala in Iceland
	BalanceAccountNumber      string
	AccountNumberPrefix       string
	AccountNumber             string
	NationalChecksum          string
	CurrencyCode              string
}

// String pretty-prints the IBAn by inserting spaces after every 4 characters.
//
// If len(IBAN)%4 != 0, then the last group will be shorter than 4 characters.
func (iban IBAN) String() string {
	var sb strings.Builder
	// maxLen plus a space for every 4 characters
	sb.Grow(maxLen + maxLen/4 + 1)

	sb.WriteString(iban.CountryCode.Code)
	sb.WriteString(twoDigitStr(int(iban.Checksum)))

	for i := 0; i < len(iban.BBAN); i += 4 {
		sb.WriteByte(' ')

		end := i + 4
		if end > len(iban.BBAN) {
			end = len(iban.BBAN)
		}

		sb.WriteString(iban.BBAN[i:end])
	}

	return sb.String()
}

func (iban IBAN) Compact() string {
	return iban.CountryCode.Code + twoDigitStr(int(iban.Checksum)) + iban.BBAN
}

var _ encoding.TextMarshaler = IBAN{}

func (iban IBAN) MarshalText() ([]byte, error) {
	return []byte(iban.Compact()), nil
}

var _ encoding.TextUnmarshaler = (*IBAN)(nil)

func (iban *IBAN) UnmarshalText(text []byte) error {
	parsed, err := Parse(string(text))
	if err != nil {
		return err
	}

	*iban = parsed
	return nil
}

func twoDigitStr(n int) string {
	return string(rune('0'+n/10)) + string(rune('0'+n%10))
}
