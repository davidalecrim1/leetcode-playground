package main

import "testing"

func TestSearchInsert(t *testing.T) {
	tests := []struct {
		scenario string
		given    []int
		target   int
		expect   int
	}{
		{
			scenario: "existing second index",
			given:    []int{1, 3, 5, 6},
			target:   5,
			expect:   2,
		},
		{
			scenario: "existing first index",
			given:    []int{1, 3, 5, 6},
			target:   1,
			expect:   0,
		},
		{
			scenario: "existing last index",
			given:    []int{1, 3, 5, 6},
			target:   7,
			expect:   4,
		},
		{
			scenario: "new in the middle index",
			given:    []int{1, 3, 5, 6},
			target:   2,
			expect:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.scenario, func(t *testing.T) {
			got := searchInsert(tt.given, tt.target)
			if got != tt.expect {
				t.Errorf("expect '%d', got '%d'", tt.expect, got)
			}
		})
	}
}
