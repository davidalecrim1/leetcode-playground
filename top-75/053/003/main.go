package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var counts [26]int
	for i := range len(s) {
		counts[s[i]-'a']++
		counts[t[i]-'a']--
	}

	for _, v := range counts {
		if v != 0 {
			return false
		}
	}

	return true
}

func main() {
	s := "david"
	r := []rune(s)
	fmt.Printf("The 'd' number: %v\n", r[0])
	fmt.Printf("The 'd' number minus 'a': %v\n", r[0]-'a')
}
