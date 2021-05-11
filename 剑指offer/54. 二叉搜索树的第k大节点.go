// -*- coding:utf-8 -*-
// @Time : 2021/5/11 11:48 上午
// @Author: MJ-CJM
// @File : leetcode/54. 二叉搜索树的第k大节点
package main

func kthLargest(root *TreeNode, k int) int {
	l := InOrderTravel_1(root)
	n := len(l)
	return l[n-k]
}

// 中序遍历
func InOrderTravel_1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	res := InOrderTravel_1(root.Left)
	res = append(res, root.Val)
	res = append(res, InOrderTravel_1(root.Right)...)
	return res
}
