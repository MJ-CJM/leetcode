// -*- coding:utf-8 -*-
// @Time : 2024/5/7 23:29
// @Author: MJ-CJM
// @File : leetcode/1474.删除链表 M 节点后的 N 个节点
package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteNodes(head *ListNode, m int, n int) *ListNode {
	k := m
	p := n
	new := head
	for new != nil {
		if k > 1 {
			k--
			new = new.Next
		} else if p > 0 {
			if new.Next == nil {
				return head
			}
			new.Next = new.Next.Next
			p--
		} else {
			new = new.Next
			k = m
			p = n
		}
	}
	return head
}
