package main

// Use a hashmap to measure the frequency
// Create a bucket sort using a matrix (slice of slice)
// Fetch using the bucket index to determine the most frequent

// Time: O(n)
// Space: O(n)
func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int, len(nums))

	for _, n := range nums {
		freq[n]++
	}

	// the bucket index is the frequency
	// meaning it's sorted by nature
	// +1 because we don't want the idx, but the actual number
	bucket := make([][]int, len(nums)+1)
	for num, total := range freq {
		// because both num 1 and 2 e.g. can have the total 3
		bucket[total] = append(bucket[total], num)
	}

	res := make([]int, 0, k)
	for i := len(bucket) - 1; i >= 0 && len(res) < k; i-- {
		res = append(res, bucket[i]...)
	}

	// safe guard if we add more then k
	return res[:k]

}
