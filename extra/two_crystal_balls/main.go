package main

import (
	"fmt"
	"math"
)

// Two Crystal Ball Problem

func FindBreakingPoint(breaks []bool) int {
	jmpAmount := int(math.Floor(math.Sqrt(float64(len(breaks)))))

	i := jmpAmount
	for ; i < len(breaks); i += jmpAmount {
		if breaks[i] {
			break
		}
	}

	breakingState := i
	i -= jmpAmount // fallback to last false state
	for ; i <= breakingState; i++ {
		if breaks[i] {
			return i
		}
	}

	return -1
}

func main() {
	// Example: Building with 100 floors, ball breaks from floor 70 onward
	building := make([]bool, 100)
	for i := 70; i < len(building); i++ {
		building[i] = true
	}

	breakingPoint := FindBreakingPoint(building)
	fmt.Println("Ball breaks at floor:", breakingPoint)
}
