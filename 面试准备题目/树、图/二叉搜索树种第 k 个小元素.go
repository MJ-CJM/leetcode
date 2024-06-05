package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil && k == 1 {
		return root.Val
	}
	stack := []*TreeNode{}
	p := root
	for p != nil || len(stack) != 0 {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}
		tmp := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]
		k--
		if k == 0 {
			return tmp.Val
		}
		p = tmp.Right
	}
	return 0
}
