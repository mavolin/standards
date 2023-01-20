package iso3166

import (
	"errors"
	"strings"
)

var ErrInvalidAlpha2 = errors.New("invalid alpha-2 code")

// ParseAlpha2 parses an ISO 3166-1 alpha-2 code.
//
// Before parsing, the code is converted to uppercase.
//
// If ParseAlpha2 returns without an error, the code is considered syntactically
// valid.
// Refer to [Alpha2Code.Status] to see if the code is in use.
func ParseAlpha2(s string) (Alpha2Code, error) {
	s = strings.ToUpper(s)

	c, ok := alpha2Codes[s]
	if !ok {
		return Alpha2Code{}, ErrInvalidAlpha2
	}

	return c, nil
}

// ValidateAlpha2 validates an ISO 3166-1 alpha-2 code, according to the
// rules laid out in [ParseAlpha2].
func ValidateAlpha2(s string) bool {
	_, err := ParseAlpha2(s)
	return err == nil
}
