package main

import "testing"

func TestStrStr(t *testing.T) {

	tests := []struct {
		scenario string
		given    string
		indexOf  string
		want     int
	}{
		{
			scenario: "first index",
			given:    "sadbutsad",
			indexOf:  "sad",
			want:     0,
		},
		{
			scenario: "last index",
			given:    "partyrock",
			indexOf:  "rock",
			want:     5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.scenario, func(t *testing.T) {
			got := strStr(tt.given, tt.indexOf)

			if got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}

}
