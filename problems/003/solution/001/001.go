package main

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for _, str := range strs[1:] {

		for len(str) < len(prefix) || !strings.HasPrefix(str, prefix) {
			prefix = prefix[:len(prefix)-1]

			if prefix == "" {
				return ""
			}
		}

	}

	return prefix
}

func main() {
	longestCommonPrefix([]string{"flight", "photo", "floor"})
}
