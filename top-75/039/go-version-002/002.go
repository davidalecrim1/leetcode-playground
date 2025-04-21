/*
This doesn't work. Only solves 50% of the cenarios.
A MinHeap is needed to control the last available room with the proper sorting.
*/

package main

import (
	"fmt"
	"sort"
)

/**
 * Definition of Interval:
 * type Interval struct {
 *     Start, End int
 * }
 */

type Interval struct {
	Start, End int
}

/**
 * @param intervals: an array of meeting time intervals
 * @return: the minimum number of conference rooms required
 */
func MinMeetingRooms(intervals []*Interval) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].Start
	})

	res := 1
	for i := 1; i < len(intervals); i++ {
		if intervals[i].Start < intervals[i-1].End {
			res++
		}
	}

	return res
}

func main() {
	res := MinMeetingRooms([]*Interval{
		{Start: 0, End: 30},
		{Start: 5, End: 10},
		{Start: 15, End: 20},
	})

	fmt.Printf("Expected 2 got %v\n", res)
}
