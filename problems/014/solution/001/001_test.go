package main

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		scenario string
		given    []string
		want     string
	}{
		{
			scenario: "prefix fl",
			given:    []string{"flower", "flow", "flight"},
			want:     "fl",
		},
		{
			scenario: "no prefix diff words",
			given:    []string{"dog", "racecar", "car"},
			want:     "",
		},
		{
			scenario: "two with same prefix, but third with not",
			given:    []string{"reflower", "flow", "flight"},
			want:     "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.scenario, func(t *testing.T) {
			got := longestCommonPrefix(tt.given)
			if got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
