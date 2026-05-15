package main

func isAnagram(s string, t string) bool {
	var seen [26]int

	if len(s) != len(t) {
		return false
	}

	for i := 0; i < len(s); i++ {
		seen[s[i]-'a']++
		seen[t[i]-'a']--
	}

	for i := range seen {
		if seen[i] != 0 {
			return false
		}
	}

	return true
}
