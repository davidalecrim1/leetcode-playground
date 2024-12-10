package main

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	t.Run("duplicate array", func(t *testing.T) {
		nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
		want := []int{0, 1, 2, 3, 4}

		got := removeDuplicates(nums)

		if got != len(want) {
			t.Errorf("got %v want %v", got, want)
		}

	})

}
