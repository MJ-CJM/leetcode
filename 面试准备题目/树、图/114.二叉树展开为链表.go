package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
	if root == nil {
		return
	}

	var head *TreeNode
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		cur := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]

		if head != nil {
			head.Right = cur
			head.Left = nil
		}

		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}

		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}

		head = cur
	}
	return
}
