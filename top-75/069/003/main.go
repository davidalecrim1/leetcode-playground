package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	track := []int{}

	var dfs func(curr *TreeNode)
	dfs = func(curr *TreeNode) {
		// left
		// root
		// right

		if curr == nil {
			return
		}

		dfs(curr.Left)
		track = append(track, curr.Val)
		dfs(curr.Right)
	}
	dfs(root)
	fmt.Println(track)

	return track[k-1]
}

// in order traversal of the bst
// store the elements in an array
// return the kth element
