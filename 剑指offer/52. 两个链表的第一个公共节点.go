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

//
