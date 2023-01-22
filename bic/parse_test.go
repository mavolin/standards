package bic

import "testing"

func TestParse(t *testing.T) {
	testCases := []struct {
		In     string
		Expect string
	}{
		{In: "BELADEBEXXX", Expect: "BELADEBEXXX"},
		{In: "RBOSGGSX", Expect: "RBOSGGSX"},
		{In: "CHASGB2LXXX", Expect: "CHASGB2LXXX"},
		{In: "RZTIAT22263", Expect: "RZTIAT22263"},
		{In: "BCEELULL", Expect: "BCEELULL"},
		{In: "MARKDEFF", Expect: "MARKDEFF"},
		{In: "GENODEF1JEV", Expect: "GENODEF1JEV"},
		{In: "UBSWCHZH80A", Expect: "UBSWCHZH80A"},
		{In: "CEDELULLXXX", Expect: "CEDELULLXXX"},
		{In: "HELADEF1RRS", Expect: "HELADEF1RRS"},
		{In: "GENODEF1S04", Expect: "GENODEF1S04"},
		{In: "BELA DEBE XXX", Expect: "BELADEBEXXX"},
		{In: "BELA    DEBE   XXX", Expect: "BELADEBEXXX"},
	}

	for _, c := range testCases {
		t.Run(c.In, func(t *testing.T) {
			bic, err := Parse(c.In)
			if err != nil {
				t.Errorf("Parse(%q): %s", c.In, err)
			}

			if bic.String() != c.Expect {
				t.Errorf("Parse(%q): expected %q, got %q", c.In, c.Expect, bic.String())
			}
		})
	}
}
