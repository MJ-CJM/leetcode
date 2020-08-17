package main

// Definition for a binary tree node.
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil{
		return [][]int{}
	}
	result := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	// visited := make([]TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		tmp := []int{}
		tmp2 := []*TreeNode{}
		for _, v := range queue{
			tmp = append(tmp, v.Val)
			if v.Left != nil {
				tmp2 = append(tmp2, v.Left)
			}
			if v.Right != nil {
				tmp2 = append(tmp2, v.Right)
			}
		}
		result = append(result, tmp)
		queue = tmp2
	}
	return result
}