// Package bic provides parsing and validation of BIC (ISO 9362) codes.
package bic

import "encoding"

// BIC is a ISO 9362 Business Identifier Code.
type BIC struct {
	// BusinessPartyPrefix, in former editions of ISO 9362 known as the
	// Bank or Institution Code, is a four-character alphanumeric code
	// that identifies the business party.
	BusinessPartyPrefix string
	// CountryCode is the two character alphabetic ISO 3166-1 alpha-2
	// country code.
	//
	// Besides the officially assigned ISO 3166-1 country codes, "XK" is used
	// to represent the Republic of Kosovo.
	CountryCode string
	// BusinessPartySuffix, in former editions of ISO 9362 known as the
	// Location Code, is a two-character alphanumeric code that identifies
	// the location of the business party.
	BusinessPartySuffix string
	// BranchCode is the three character optional alphanumeric branch code.
	//
	// Some BICs may use "XXX" as branch code, which is equivalent to an empty
	// branch code and denotes the primary office of the institution.
	BranchCode string
}

func (bic BIC) String() string {
	return bic.BusinessPartyPrefix + bic.CountryCode + bic.BusinessPartySuffix + bic.BranchCode
}

var _ encoding.TextMarshaler = BIC{}

func (bic BIC) MarshalText() ([]byte, error) {
	return []byte(bic.String()), nil
}

var _ encoding.TextUnmarshaler = (*BIC)(nil)

func (bic *BIC) UnmarshalText(text []byte) error {
	b, err := Parse(string(text))
	if err != nil {
		return err
	}

	*bic = b
	return nil
}
