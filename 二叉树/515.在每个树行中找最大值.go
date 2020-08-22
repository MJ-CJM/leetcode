package main

import "math"

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	res := []int{}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		max := math.MinInt32
		tmp2 := []*TreeNode{}
		for _, v := range queue {
			if v.Val > max {
				max = v.Val
			}
			if v.Left != nil {
				tmp2 = append(tmp2, v.Left)
			}
			if v.Right != nil {
				tmp2 = append(tmp2, v.Right)
			}
		}
		res = append(res, max)
		queue = tmp2
	}
	return res
}
