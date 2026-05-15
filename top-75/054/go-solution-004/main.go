package main

func groupAnagrams(strs []string) [][]string {
	// Strategy
	// Sort the string
	// Use a hashmap to group them

	// Use a fixed array 26 to count the letters
	// Use the array as key to the hashmap

	hm := make(map[[26]int][]string, len(strs))
	for _, str := range strs {
		idx := [26]int{}
		for i := range str {
			idx[str[i]-'a']++
		}

		hm[idx] = append(hm[idx], str)
	}

	res := make([][]string, 0, len(hm))
	for _, v := range hm {
		res = append(res, v)
	}

	return res
}
