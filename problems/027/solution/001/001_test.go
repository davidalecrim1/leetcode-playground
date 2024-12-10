package main

import (
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {

	assertExpectedSlice := func(t testing.TB, nums []int, got int, expect []int) {
		if !reflect.DeepEqual(expect, nums[:got]) {
			t.Errorf("got %v, expect %v", nums[:got], expect)
		}
	}

	t.Run("remove val 3", func(t *testing.T) {
		nums := []int{3, 2, 2, 3}
		got := RemoveElement(nums, 3)
		assertExpectedSlice(t, nums, got, []int{2, 2})
	})

	t.Run("remove val 2", func(t *testing.T) {
		nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
		got := RemoveElement(nums, 2)
		assertExpectedSlice(t, nums, got, []int{0, 1, 3, 0, 4})
	})
}
