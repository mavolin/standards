package postalcode

import "testing"

func TestParse(t *testing.T) {
	testCases := []struct {
		In     string
		Expect string
	}{
		{In: "26123", Expect: "26123"},
		{In: "21149", Expect: "21149"},
	}

	for _, c := range testCases {
		t.Run(c.In, func(t *testing.T) {
			pin, err := Parse(c.In)
			if err != nil {
				t.Errorf("Parse(%q): %s", c, err)
			}

			if pin.String() != c.Expect {
				t.Errorf("Parse(%q): expected %q, got %q", c.In, c.Expect, pin.String())
			}
		})
	}
}
