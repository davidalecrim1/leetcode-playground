package main

func characterReplacement(s string, k int) int {
	l := 0
	seen := [26]int{}
	res := 0

	for r := 0; r < len(s); r++ {
		// keep track of the chars
		seen[s[r]-'A']++

		longest := 0
		for i := range seen {
			longest = max(longest, seen[i])
		}

		// make the window valid again
		for longest+k < (r-l)+1 {
			seen[s[l]-'A'] -= 1
			l++
		}

		res = max(res, (r-l)+1)
	}

	return res
}
