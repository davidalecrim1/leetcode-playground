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

// Time Complexity: O(n^2)
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) < 1 {
		return nil
	}

	head := &TreeNode{Val: preorder[0]}

	rootIdxInOrder := -1
	for i, val := range inorder {
		if val == preorder[0] {
			rootIdxInOrder = i
			break
		}
	}
	leftSize := rootIdxInOrder

	head.Left = buildTree(preorder[1:leftSize+1], inorder[:rootIdxInOrder])
	head.Right = buildTree(preorder[leftSize+1:], inorder[rootIdxInOrder+1:])

	return head
}

func main() {
	fmt.Printf("buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}): %v\n", buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
}
