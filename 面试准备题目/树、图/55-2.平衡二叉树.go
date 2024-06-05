package main

import "math"

func isBalanced(root *TreeNode) bool {
	return root == nil || (isBalanced(root.Left) && isBalanced(root.Right) && math.Abs(getHeight(root.Left)-getHeight(root.Right)) < 2)
}

func getHeight(root *TreeNode) float64 {
	if root == nil {
		return 0
	}
	left := getHeight(root.Left)
	right := getHeight(root.Right)
	return math.Max(left, right) + 1
}