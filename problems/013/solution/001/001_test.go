package main

import "testing"

func TestRomanToInt(t *testing.T) {

	tests := []struct {
		scenario string
		given    string
		expect   int
	}{
		{
			scenario: "III",
			given:    "III",
			expect:   3,
		},
		{
			scenario: "LVIII",
			given:    "LVIII",
			expect:   58,
		},
		{
			scenario: "XXIX",
			given:    "XXIX",
			expect:   29,
		},
		{
			scenario: "XXXIV",
			given:    "XXXIV",
			expect:   34,
		},
		{
			scenario: "XLIX",
			given:    "XLIX",
			expect:   49,
		},
		{
			scenario: "XCIX",
			given:    "XCIX",
			expect:   99,
		},
		{
			scenario: "CXII",
			given:    "CXII",
			expect:   112,
		},
		{
			scenario: "CDXLIV",
			given:    "CDXLIV",
			expect:   444,
		},
		{
			scenario: "CMXCIV",
			given:    "CMXCIV",
			expect:   994,
		},
		{
			scenario: "MCMXCIV",
			given:    "MCMXCIV",
			expect:   1994,
		},
	}

	for _, test := range tests {

		executeTest := func(t *testing.T) {
			got := RomanToInt(test.given)

			if got != test.expect {
				t.Errorf("expected %v, got %v", test.expect, got)
			}
		}

		t.Run(test.scenario, executeTest)
	}

}

func BenchmarkRomanToInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = RomanToInt("MCMXCIV")
	}
}
