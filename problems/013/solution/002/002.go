package main

import "fmt"

var romanValues = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func RomanToInt(s string) int {
	total := 0
	lastValue := 0

	for i := len(s) - 1; i >= 0; i-- {
		currentValue := romanValues[s[i]]

		if currentValue < lastValue {
			total -= currentValue
		} else {
			total += currentValue
		}
		lastValue = currentValue
	}

	return total
}

func main() {
	fmt.Println(RomanToInt("MCMXCIV"))
}
