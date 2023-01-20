// Package validate provides validation utilities.
package validate

// ============================================================================
// Alphabetic
// ======================================================================================

// ===================================== Lowercase ======================================

func IsLowerAlpha(s string) bool {
	for _, r := range s {
		if !IsRuneLowerAlpha(r) {
			return false
		}
	}

	return true
}

func IsRuneLowerAlpha(r rune) bool {
	return r >= 'a' && r <= 'z'
}

// ===================================== Uppercase ======================================

func IsUpperAlpha(s string) bool {
	for _, r := range s {
		if !IsRuneUpperAlpha(r) {
			return false
		}
	}

	return true
}

func IsRuneUpperAlpha(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

// ================================== Case Insensitive ==================================

func IsAlpha(s string) bool {
	for _, r := range s {
		if !IsRuneAlpha(r) {
			return false
		}
	}

	return true
}

func IsRuneAlpha(r rune) bool {
	return IsRuneLowerAlpha(r) || IsRuneUpperAlpha(r)
}

// ============================================================================
// Numeric
// ======================================================================================

func IsNumeric(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}

	return true
}

func IsRuneNumeric(r rune) bool {
	return r >= '0' && r <= '9'
}

// ============================================================================
// Alphanumeric
// ======================================================================================

// ===================================== Lowercase ======================================

func IsLowerAlphanumeric(s string) bool {
	for _, r := range s {
		if !IsRuneLowerAlphanumeric(r) {
			return false
		}
	}

	return true
}

func IsRuneLowerAlphanumeric(r rune) bool {
	return IsRuneLowerAlpha(r) || IsRuneNumeric(r)
}

// ===================================== Uppercase ======================================

func IsUpperAlphanumeric(s string) bool {
	for _, r := range s {
		if !IsRuneUpperAlphanumeric(r) {
			return false
		}
	}

	return true
}

func IsRuneUpperAlphanumeric(r rune) bool {
	return IsRuneUpperAlpha(r) || IsRuneNumeric(r)
}

// ================================== Case Insensitive ==================================

func IsAlphanumeric(s string) bool {
	for _, r := range s {
		if !IsRuneAlphanumeric(r) {
			return false
		}
	}

	return true
}

func IsRuneAlphanumeric(r rune) bool {
	return IsRuneAlpha(r) || IsRuneNumeric(r)
}
