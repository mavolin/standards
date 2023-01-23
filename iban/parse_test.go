package iban

import "testing"

func TestParse(t *testing.T) {
	// https://ibanvalidieren.de/beispiele.html
	// 2023-01-22
	testCases := []struct {
		In     string
		Expect string
	}{
		{In: "DE02120300000000202051", Expect: "DE02 1203 0000 0000 2020 51"},
		{In: "DE02500105170137075030", Expect: "DE02 5001 0517 0137 0750 30"},
		{In: "DE88100900001234567892", Expect: "DE88 1009 0000 1234 5678 92"},
		{In: "AT026000000001349870", Expect: "AT02 6000 0000 0134 9870"},
		{In: "AT021420020010147558", Expect: "AT02 1420 0200 1014 7558"},
		{In: "AT023200000000641605", Expect: "AT02 3200 0000 0064 1605"},
		{In: "CH0209000000100013997", Expect: "CH02 0900 0000 1000 1399 7"},
		{In: "CH0204835000626882001", Expect: "CH02 0483 5000 6268 8200 1"},
		{In: "CH0200700110000387896", Expect: "CH02 0070 0110 0003 8789 6"},
		{In: "LI0208800000017197386", Expect: "LI02 0880 0000 0171 9738 6"},
		{In: "LI0508812105028570001", Expect: "LI05 0881 2105 0285 7000 1"},
		{In: "LI2608802001003488101", Expect: "LI26 0880 2001 0034 8810 1"},

		{In: "LI26 0 8802 0010 0348 8101", Expect: "LI26 0880 2001 0034 8810 1"},
		{In: "LI26 0  88 02  00 10  0348  8 101", Expect: "LI26 0880 2001 0034 8810 1"},

		// National Checksums

		{In: "AL47 2121 1009 0000 0002 3569 8741", Expect: "AL47 2121 1009 0000 0002 3569 8741"},
		{In: "BE68 5390 0754 7034", Expect: "BE68 5390 0754 7034"},
		{In: "BA39 1290 0794 0102 8494", Expect: "BA39 1290 0794 0102 8494"},
		{In: "HR12 1001 0051 8630 0016 0", Expect: "HR12 1001 0051 8630 0016 0"},
		{In: "CZ65 0800 0000 1920 0014 5399", Expect: "CZ65 0800 0000 1920 0014 5399"},
		{In: "TL38 0080 0123 4567 8910 157", Expect: "TL38 0080 0123 4567 8910 157"},
		{In: "EE38 2200 2210 2014 5685", Expect: "EE38 2200 2210 2014 5685"},
		{In: "FI21 1234 5600 0007 85", Expect: "FI21 1234 5600 0007 85"},
		{In: "FR14 2004 1010 0505 0001 3M02 606", Expect: "FR14 2004 1010 0505 0001 3M02 606"},
		{In: "HU42 1177 3016 1111 1018 0000 0000", Expect: "HU42 1177 3016 1111 1018 0000 0000"},
		{In: "IS14 0159 2600 7654 5510 7303 39", Expect: "IS14 0159 2600 7654 5510 7303 39"},
		{In: "MK07 2501 2000 0058 984", Expect: "MK07 2501 2000 0058 984"},
		{In: "ME25 5050 0001 2345 6789 51", Expect: "ME25 5050 0001 2345 6789 51"},
		{In: "NO93 8601 1117 947", Expect: "NO93 8601 1117 947"},
		{In: "PL61 1090 1014 0000 0712 1981 2874", Expect: "PL61 1090 1014 0000 0712 1981 2874"},
		{In: "MC58 1122 2000 0101 2345 6789 030", Expect: "MC58 1122 2000 0101 2345 6789 030"},
		{In: "PT50 0002 0123 1234 5678 9015 4", Expect: "PT50 0002 0123 1234 5678 9015 4"},
		{In: "RS352 600 0560 1001 6113 79", Expect: "RS35 2600 0560 1001 6113 79"},
		{In: "SK31 1200 0000 1987 4263 7541", Expect: "SK31 1200 0000 1987 4263 7541"},
		{In: "ES91 2100 0418 4502 0005 1332", Expect: "ES91 2100 0418 4502 0005 1332"},
		{In: "SI56 2633 0001 2039 086", Expect: "SI56 2633 0001 2039 086"},
	}

	for _, c := range testCases {
		t.Run(c.In, func(t *testing.T) {
			iban, err := Parse(c.In)
			if err != nil {
				t.Errorf("Parse(%q): %s", c.In, err)
			}

			if iban.String() != c.Expect {
				t.Errorf("Parse(%q): expected %q, got %q", c.In, c.Expect, iban.String())
			}
		})
	}
}