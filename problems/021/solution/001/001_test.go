package main

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		scenario   string
		givenList1 *ListNode
		givenList2 *ListNode
		want       *ListNode
	}{
		{
			scenario:   "Both lists are empty",
			givenList1: nil,
			givenList2: nil,
			want:       nil,
		},
		{
			scenario:   "First list is empty, second is non-empty",
			givenList1: nil,
			givenList2: &ListNode{1, &ListNode{3, &ListNode{5, nil}}},
			want:       &ListNode{1, &ListNode{3, &ListNode{5, nil}}},
		},
		{
			scenario:   "First list is non-empty, second is empty",
			givenList1: &ListNode{2, &ListNode{4, &ListNode{6, nil}}},
			givenList2: nil,
			want:       &ListNode{2, &ListNode{4, &ListNode{6, nil}}},
		},
		{
			scenario:   "Both lists are non-empty",
			givenList1: &ListNode{1, &ListNode{3, &ListNode{5, nil}}},
			givenList2: &ListNode{2, &ListNode{4, &ListNode{6, nil}}},
			want:       &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}},
		},
		{
			scenario:   "Similar lists",
			givenList1: &ListNode{1, &ListNode{2, &ListNode{4, nil}}},
			givenList2: &ListNode{1, &ListNode{3, &ListNode{4, nil}}},
			want:       &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, nil}}}}}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.scenario, func(t *testing.T) {
			got := mergeTwoLists(tt.givenList1, tt.givenList2)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
