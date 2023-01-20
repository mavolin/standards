package iso3166

import "encoding"

//go:generate go run github.com/mavolin/miscnums/tools/codegen/iso3166-1

// ============================================================================
// Alpha2Code
// ======================================================================================

type Alpha2Code struct {
	// Code is the uppercase ISO 3166-1 alpha-2 code.
	Code string
	// Status is the status of the code.
	Status Status
	// Country is the name of the country it belongs to.
	Country string
}

func (c Alpha2Code) String() string {
	return c.Code
}

var _ encoding.TextMarshaler = Alpha2Code{}

func (c Alpha2Code) MarshalText() ([]byte, error) {
	return []byte(c.Code), nil
}

var _ encoding.TextUnmarshaler = (*Alpha2Code)(nil)

func (c *Alpha2Code) UnmarshalText(text []byte) error {
	code, err := ParseAlpha2(string(text))
	if err != nil {
		return err
	}

	*c = code
	return nil
}

// ============================================================================
// Status
// ======================================================================================

type Status uint8

const (
	OfficiallyAssigned Status = iota + 1
	UserAssigned
	ExceptionallyReserved
	TransitionallyReserved
	IndeterminatelyReserved
	FormerlyAssigned
	Unassigned
)

func (s Status) IsAssigned() bool {
	return s == OfficiallyAssigned || s == UserAssigned
}

func (s Status) IsReserved() bool {
	return s == ExceptionallyReserved || s == TransitionallyReserved || s == IndeterminatelyReserved
}

func (s Status) String() string {
	switch s {
	case OfficiallyAssigned:
		return "officially assigned"
	case UserAssigned:
		return "user assigned"
	case ExceptionallyReserved:
		return "exceptionally reserved"
	case TransitionallyReserved:
		return "transitionally reserved"
	case IndeterminatelyReserved:
		return "indeterminately reserved"
	case FormerlyAssigned:
		return "formerly assigned"
	case Unassigned:
		return "unassigned"
	default:
		return "unknown"
	}
}
