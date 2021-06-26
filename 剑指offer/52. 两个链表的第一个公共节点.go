// -*- coding:utf-8 -*-
// @Time : 2021/5/11 10:29 上午
// @Author: MJ-CJM
// @File : leetcode/52. 两个链表的第一个公共节点
package main

// 利用字典的特性来判断
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	mapList1 := make(map[*ListNode]*ListNode)

	for headA != nil {
		mapList1[headA] = headA
		headA = headA.Next
	}
	for headB != nil {
		if mapList1[headB] == headB {
			return headB
		}
		headB = headB.Next
	}
	return nil
}

// 原地循环等待
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	node1, node2 := headA, headB
	for node1 != node2 {
		if node1 != nil {
			node1 = node1.Next
		} else {
			node1 = headB
		}
		if node2 != nil {
			node2 = node2.Next
		} else {
			node2 = headA
		}
	}
	return node1
}
