// -*- coding:utf-8 -*-
// @Time : 2024/5/10 00:13
// @Author: MJ-CJM
// @File : leetcode/2. 两数相加
package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sumTmp := 0
	newNode := &ListNode{}
	head := newNode

	flag := 0
	for l1 != nil && l2 != nil {
		if flag == 1 {
			sumTmp = l1.Val + l2.Val + 1
			flag = 0
		} else {
			sumTmp = l1.Val + l2.Val
		}
		if sumTmp <= 9 {
			nextNode := ListNode{
				Val: sumTmp,
			}
			newNode.Next = &nextNode
		} else {
			flag = 1
			tmpCount := sumTmp % 10
			nextNode := ListNode{
				Val: tmpCount,
			}
			newNode.Next = &nextNode
		}
		newNode = newNode.Next
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		if flag == 1 {
			sumTmp := l1.Val + 1
			if sumTmp <= 9 {
				flag = 0
				nextNode := ListNode{
					Val: sumTmp,
				}
				newNode.Next = &nextNode
			} else {
				tmpCount := sumTmp % 10
				nextNode := ListNode{
					Val: tmpCount,
				}
				newNode.Next = &nextNode
			}
		} else {
			nextNode := ListNode{
				Val: l1.Val,
			}
			newNode.Next = &nextNode
		}
		newNode = newNode.Next
		l1 = l1.Next
	}

	for l2 != nil {
		if flag == 1 {
			sumTmp := l2.Val + 1
			if sumTmp <= 9 {
				flag = 0
				nextNode := ListNode{
					Val: sumTmp,
				}
				newNode.Next = &nextNode
			} else {
				tmpCount := sumTmp % 10
				nextNode := ListNode{
					Val: tmpCount,
				}
				newNode.Next = &nextNode
			}
		} else {
			nextNode := ListNode{
				Val: l2.Val,
			}
			newNode.Next = &nextNode
		}
		newNode = newNode.Next
		l2 = l2.Next
	}

	if flag == 1 {
		nextNode := ListNode{
			Val: 1,
		}
		newNode.Next = &nextNode
	}

	return head.Next
}
