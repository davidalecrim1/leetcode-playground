package main

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

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}

	rootIdxInOrder := 0
	for i, val := range inorder {
		if val == root.Val {
			rootIdxInOrder = i
			break
		}
	}

	leftSize := rootIdxInOrder

	root.Left = buildTree(preorder[1:1+leftSize], inorder[:rootIdxInOrder])
	root.Right = buildTree(preorder[1+leftSize:], inorder[rootIdxInOrder+1:])

	return root
}
