package main

import "slices"

// use a sliding window the size of s1 in s2
// when found caracter that contains in s1 on s2, sort the window and the s1
// then compare them if they are equal

func checkInclusion(s1 string, s2 string) bool {
	b1 := createBucket([]byte(s1))

	l := 0
	r := len(s1)

	ss2 := []byte(s2)
	for r <= len(s2) {
		b2 := createBucket(ss2[l:r])
		if slices.Equal(b1[:], b2[:]) {
			return true
		}

		l++
		r++
	}

	return false
}

func createBucket(in []byte) [26]byte {
	out := [26]byte{}
	for i := range in {
		out[in[i]-'a']++
	}
	return out
}
