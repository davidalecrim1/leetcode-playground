package main

// if num[i-1] is == num[i], not consecutive
// if num[i-1]+1 is not == num[i], not consecutive

// Sorting will make this O(n log n)

// Make the nums into a set
// Identify the start of a sequence with a num where num[i-1] doesn't exist
// Walk the sequence and count the window, store in a var with max

func longestConsecutive(nums []int) int {
	set := make(map[int]struct{})

	// make the nums a set
	for _, n := range nums {
		set[n] = struct{}{}
	}

	longest := 0
	for n, _ := range set {
		// not the start of a sequence
		if _, ok := set[n-1]; ok {
			continue
		}

		sequence := 0
		curr := n

	inner:
		for {
			// if the next is not found, break
			if _, ok := set[curr]; !ok {
				break inner
			}

			sequence++
			curr++
			longest = max(longest, sequence)
		}
	}

	return longest
}
