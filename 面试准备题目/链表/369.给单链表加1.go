// -*- coding:utf-8 -*-
// @Time : 2024/5/8 01:35
// @Author: MJ-CJM
// @File : leetcode/369.给单链表加1
package main

// 迭代方式反转单链表
func reverseListIteratively(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	// 遍历并改变指针方向
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}


func plusOne(head *ListNode) *ListNode {
	// 先反转链表
	var pre *ListNode

	pre = reverseListIteratively(head)

	start := pre
	flag := 0
	end := pre
	for pre != nil {
		tmp := pre.Val + 1
		if tmp <= 9 {
			pre.Val++
			flag = 0
			break
		} else {
			pre.Val = 0
			flag = 1
		}
		end = pre
		pre = pre.Next
	}
	if flag == 1 {
		end.Next = &ListNode{
			Val: 1,
			Next: nil,
		}
	}

	res := reverseListIteratively(start)
	return res
}
