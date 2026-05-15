package main

func groupAnagrams(strs []string) [][]string {
	/*
	   Strategy to group the strs together
	   Use a hashmap to track the pairs.

	   Op1 Sort the string and use an index in the map
	   Walk the map and create the result
	*/

	hm := make(map[[26]byte][]string, len(strs))
	for _, str := range strs {
		var key [26]byte
		for i := range str {
			key[str[i]-'a']++
		}

		hm[key] = append(hm[key], str)
	}

	res := make([][]string, 0, len(hm))
	for _, v := range hm {
		res = append(res, v)
	}

	return res
}
