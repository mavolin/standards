package bic

import (
	"errors"
	"strings"

	"github.com/mavolin/standards/internal/validate"
	"github.com/mavolin/standards/iso3166"
)

var (
	ErrLength                   = errors.New("bic: invalid length")
	ErrCountryCodeInappropriate = errors.New("bic: country code is neither officially assigned nor user-assigned")
	ErrBusinessPartyPrefix      = errors.New("bic: invalid business party prefix")
	ErrBusinessPartySuffix      = errors.New("bic: invalid business party suffix")
	ErrBranchCode               = errors.New("bic: invalid branch code")
)

// Parse parses s into a [BIC].
//
// If Parse returns without an error, the BIC is considered syntactically
// correct.
//
// # Validation Notes
//
// Despite not formally allowed by ISO 9362, Parse allows spaces and lowercase
// letters in s.
// The individual parts of the returned BIC are always uppercase and contain no
// spaces.
//
// While SWIFT has, and will for the foreseeable future, only assigned
// letters to the business party prefix (c.f.
// [1]: Entity Identifiers/BIC Structure and [2]: section 4.3), ISO 9362
// expressly allows digits in the business party prefix.
// Therefore, Parse DOES allow digits in the business party prefix.
//
// [1]: https://www2.swift.com/knowledgecentre/publications/bic_policy
// [2]: https://web.archive.org/web/20160629235645/https://www.swift.com/sites/default/files/resources/swift_standards_whitepaper_iso93622014bicimplementation.pdf
func Parse(s string) (BIC, error) {
	s = strings.ReplaceAll(s, " ", "")
	s = strings.ToUpper(s)

	if len(s) != 8 && len(s) != 11 {
		return BIC{}, ErrLength
	}

	bic := BIC{
		BusinessPartyPrefix: s[:4],
		CountryCode:         s[4:6],
		BusinessPartySuffix: s[6:8],
		BranchCode:          s[8:],
	}

	if !validate.IsUpperAlphanumeric(bic.BusinessPartyPrefix) {
		return BIC{}, ErrBusinessPartyPrefix
	}

	cc, err := iso3166.ParseAlpha2(bic.CountryCode)
	if err != nil {
		return BIC{}, err
	}
	if !cc.Status.IsAssigned() {
		return BIC{}, ErrCountryCodeInappropriate
	}

	if !validate.IsUpperAlphanumeric(bic.BusinessPartySuffix) {
		return BIC{}, ErrBusinessPartySuffix
	}

	if !validate.IsUpperAlphanumeric(bic.BranchCode) {
		return BIC{}, ErrBranchCode
	}

	return bic, nil
}

// IsValid checks whether s is a syntactically correct BIC, as defined in
// [Parse].
func IsValid(s string) bool {
	_, err := Parse(s)
	return err == nil
}
