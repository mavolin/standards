package bic

import "testing"

func TestValidate(t *testing.T) {
	testCases := []string{
		"BELADEBEXXX", "RBOSGGSX", "CHASGB2LXXX", "RZTIAT22263", "BCEELULL",
		"MARKDEFF", "GENODEF1JEV", "UBSWCHZH80A", "CEDELULLXXX", "HELADEF1RRS",
		"GENODEF1S04",
	}

	for _, c := range testCases {
		t.Run(c, func(t *testing.T) {
			if !IsValid(c) {
				t.Errorf("expected %q to be valid", c)
			}
		})
	}
}
