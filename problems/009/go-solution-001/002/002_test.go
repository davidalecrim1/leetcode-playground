package main

import "testing"

func TestIsPalindrome(t *testing.T) {

	tests := []struct {
		scenario string
		input    int
		want     bool
	}{
		{
			scenario: "negative number invalid",
			input:    -123321,
			want:     false,
		},
		{
			scenario: "zero number invalid",
			input:    0,
			want:     false,
		},
		{
			scenario: "two digit invalid",
			input:    10,
			want:     false,
		},
		{
			scenario: "two digit valid",
			input:    11,
			want:     true,
		},
		{
			scenario: "three digit invalid",
			input:    123,
			want:     false,
		},
		{
			scenario: "three digit valid",
			input:    121,
			want:     true,
		},
		{
			scenario: "seven digit valid",
			input:    1237321,
			want:     true,
		},
		{
			scenario: "seven digit invalid",
			input:    1237331,
			want:     false,
		},
	}

	for _, tt := range tests {
		executeTest := func(t *testing.T) {
			got := isPalindrome(tt.input)

			if got != tt.want {
				t.Errorf("expected %v, got %v", tt.want, got)
			}
		}

		t.Run(tt.scenario, executeTest)
	}
}
