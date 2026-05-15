package main

func groupAnagrams(strs []string) [][]string {
	groups := make(map[[26]int][]string)

	for i := range strs {
		key := [26]int{}
		str := strs[i]
		for j := range str {
			key[str[j]-'a']++
		}
		groups[key] = append(groups[key], str)
	}

	res := make([][]string, 0, len(strs))
	for _, v := range groups {
		res = append(res, v)
	}

	return res
}
