package main

// Goal is to reach N
// N can be reached with 1 step
// or with 2 step.

// This will result in repetable problems.
// Lets say N is 10
// I'm in the house 8.
// I have two ways to each the top.
// go two by one.
// go two by two.
// threfore we have subproblems that we can reuse.

// Time Complexity: O(2^N)
// Because we are doing repetetive work (in a computational way).
// Following two paths (n-1) and (n-2).
func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}

	return climbStairs(n-1) + climbStairs(n-2)
}
