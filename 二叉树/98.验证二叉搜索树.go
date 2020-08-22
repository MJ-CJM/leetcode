package main

import "math"

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	stack := []*TreeNode{}
	tmp := math.MinInt64
	p := root
	for p != nil || len(stack) != 0 {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}
		use := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if use.Val > tmp {
			tmp = use.Val
		}else{
			return false
		}
		p = use.Right
	}
	return true
}