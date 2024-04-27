// -*- coding:utf-8 -*-
// @Time : 2022/7/31 11:37 PM
// @Author: MJ-CJM
// @File : leetcode/2022.7.31-1161
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//由于计算的是每层的元素之和，用广度优先搜索来遍历这棵树会更加自然。
//对于广度优先搜索，我们可以用队列来实现。初始时，队列只包含根节点；然后不断出队，将子节点入队，直到队列为空。

func maxLevelSum(root *TreeNode) int {
	ans, maxSum := 1, root.Val
	queue := []*TreeNode{root}
	for level := 1; len(queue) > 0; level++ {
		tmp := queue
		queue = nil
		sum := 0
		for _, node := range tmp {
			sum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if sum > maxSum {
			ans, maxSum = level, sum
		}
	}
	return ans
}
