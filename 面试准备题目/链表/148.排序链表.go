package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 采用归并排序，时间 O（logn）,空间：O(1)
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 用快慢指针把链表分为两部分
	slow, fast := head, head
	pre := &ListNode{}
	for fast != nil && fast.Next != nil {
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	pre.Next = nil

	left := sortList(head)
	right := sortList(slow)
	return mergeList(left, right)
}

func mergeList(left, right *ListNode) *ListNode {
	head := &ListNode{}
	current := head

	for left != nil && right != nil {
		if left.Val < right.Val {
			current.Next = left
			left = left.Next
		} else {
			current.Next = right
			right = right.Next
		}
		current = current.Next
	}

	if left != nil {
		current.Next = left
	}

	if right != nil {
		current.Next = right
	}

	return head.Next
}
