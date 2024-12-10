package main

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		scenario string
		given    string
		want     bool
	}{
		{
			scenario: "valid open brakets",
			given:    "()",
			want:     true,
		},
		{
			scenario: "valid multiple open",
			given:    "()[]{}",
			want:     true,
		},
		{
			scenario: "invalid open brackets",
			given:    "(]",
			want:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.scenario, func(t *testing.T) {
			got := IsValid(tt.given)
			if got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
