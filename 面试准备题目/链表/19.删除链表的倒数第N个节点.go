package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	node := &ListNode{0, head}
	first := node
	second := node
	for i := 0; i <= n; i++ {
		second = second.Next
	}
	for second != nil {
		first = first.Next
		second = second.Next
	}
	first.Next = first.Next.Next
	return node.Next
}
