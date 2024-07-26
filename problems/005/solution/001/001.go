package main

import "strings"

func RomanToInt(s string) int {

	amountCM := strings.Count(s, "CM")
	if amountCM > 0 {
		s = strings.Replace(s, "CM", "", -1)
	}

	amountCD := strings.Count(s, "CD")
	if amountCD > 0 {
		s = strings.Replace(s, "CD", "", -1)
	}

	amountXC := strings.Count(s, "XC")
	if amountXC > 0 {
		s = strings.Replace(s, "XC", "", -1)
	}

	amountXL := strings.Count(s, "XL")
	if amountXL > 0 {
		s = strings.Replace(s, "XL", "", -1)
	}

	amountIV := strings.Count(s, "IV")
	if amountIV > 0 {
		s = strings.Replace(s, "IV", "", -1)
	}

	amountIX := strings.Count(s, "IX")
	if amountIX > 0 {
		s = strings.Replace(s, "IX", "", -1)
	}

	amountM := strings.Count(s, "M")
	amountC := strings.Count(s, "C")
	amountX := strings.Count(s, "X")
	amountL := strings.Count(s, "L")
	amountV := strings.Count(s, "V")
	amountI := strings.Count(s, "I")

	return amountM*1000 + amountCM*900 + amountCD*400 + amountC*100 + amountXC*90 + amountL*50 + amountXL*40 + amountX*10 + amountV*5 + amountIX*9 + amountIV*4 + amountI
}
