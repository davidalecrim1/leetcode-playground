package main

func characterReplacement(s string, k int) int {
	/*
	   Time Complexity: O(n)
	   Space Complexity: O(1)

	   Strategy:
	   - Use a sliding window to keep track of the longest substring with at most k replacements.
	   - Use a frequency array to keep track of the frequency of each character in the current window.
	   - Use a max frequency variable to keep track of the maximum frequency of any character in the current window.
	   - If the length of the current window minus the max frequency is greater than k, we need to shrink the window from the left.
	   - Update the result with the maximum of the current result and the length of the current window.
	*/

	arr := [26]int{}
	l := 0
	res := 0

	for r := range s {
		arr[int(s[r]-'A')]++

		maxFreq := maxf(arr)
		for (r-l+1)-maxFreq > k {
			arr[int(s[l]-'A')]--
			l++
		}

		res = max(res, r-l+1)
	}

	return res
}

func maxf(arr [26]int) int {
	m := arr[0]
	for i := range arr {
		m = max(m, arr[i])
	}
	return m
}
