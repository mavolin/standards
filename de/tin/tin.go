// Package tin provides parsing and validation for German Tax Identification
// Numbers (Steuerliche Identifikationsnummer).
package tin

import (
	"encoding"
	"strconv"
)

// TIN is a German Tax Identification Number (Steuerliche
// Identifikationsnummer).
//
// It does not contain any spaces.
type TIN uint64

// String returns a human-readable representation of the tax ID, by separating
// the tax ID into its parts.
//
// The returned string is in the format "XX XXX XXX XXX".
//
// If the TIN is invalid, it is returned as is.
//
// It is guaranteed that for any valid TIN, the returned string can be parsed
// back into the same TIN.
func (id TIN) String() string {
	if id < 10_000_000_000 || id > 99_999_999_999 {
		return strconv.FormatUint(uint64(id), 10)
	}

	s := strconv.FormatUint(uint64(id), 10)
	return s[:2] + " " + s[2:5] + " " + s[5:8] + " " + s[8:]
}

var _ encoding.TextMarshaler = TIN(0)

// MarshalText implements encoding.TextMarshaler.
//
// In contrast to String(), it returns the TIN as is, without any formatting.
func (id TIN) MarshalText() ([]byte, error) {
	s := strconv.FormatUint(uint64(id), 10)
	return []byte(s), nil
}

var _ encoding.TextUnmarshaler = (*TIN)(nil)

func (id *TIN) UnmarshalText(text []byte) error {
	parsedID, err := Parse(string(text))
	if err != nil {
		return err
	}

	*id = parsedID
	return nil
}
