package main

import "fmt"

// Use a module to keep track of the clockwise circle
// Keep track of the index to move in the circle

func findTheWinner(n int, k int) int {
	idx := 0

	friends := make([]int, n)
	for i := range n {
		friends[i] = i + 1
	}

	for len(friends) != 1 {
		total := len(friends)
		idx = (idx + k - 1) % total
		friends = append(friends[:idx], friends[idx+1:]...)
		fmt.Println(idx)
	}

	return friends[0]
}
