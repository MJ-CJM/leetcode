// -*- coding:utf-8 -*-
// @Time : 2024/5/8 00:13
// @Author: MJ-CJM
// @File : leetcode/708.循环有序列表的插入
package main

type Node struct {
	Val  int
	Next *Node
}

func insert(aNode *Node, x int) *Node {
	xNode := &Node{
		Val: x,
	}

	if aNode == nil {
		xNode.Next = xNode
		return xNode
	}

	cur := aNode
	for {
		// 查找合适的点
		if cur.Val <= x && cur.Next.Val >= x {
			break
		}
		// 查找收尾交替的情况
		if cur.Val > cur.Next.Val {
			if x >= cur.Val || x <= cur.Next.Val {
				break
			}
		}
		// 如果已经回到起点
		if cur.Next == aNode {
			break
		}
		cur = cur.Next
	}

	xNode.Next = cur.Next
	cur.Next = xNode

	return aNode
}