// -*- coding:utf-8 -*-
// @Time : 2022/8/17 11:30 PM
// @Author: MJ-CJM
// @File : leetcode/2022.8.17-1302. 层数最深叶子节点的和
package main

func deepestLeavesSum(root *TreeNode) int {
	var result int
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		tmp := 0
		tmp2 := make([]*TreeNode, 0)
		for _, v := range queue {
			tmp += v.Val
			if v.Left != nil {
				tmp2 = append(tmp2, v.Left)
			}
			if v.Right != nil {
				tmp2 = append(tmp2, v.Right)
			}
		}
		queue = tmp2
		result = tmp
	}
	return result
}
