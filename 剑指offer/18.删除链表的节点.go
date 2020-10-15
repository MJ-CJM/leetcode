package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	pre := &ListNode{
		Next: head,
	}
	res := pre
	cur := head
	for cur != nil {
		if cur.Val == val {
			pre.Next = pre.Next.Next
			return res.Next
		}
		cur = cur.Next
		pre = pre.Next
	}
	return nil
}
