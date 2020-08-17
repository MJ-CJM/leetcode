package main


func reverseKGroup(head *ListNode, k int) *ListNode {
	first := &ListNode{Next: head}
	pre := first
	cur := head
	for cur != nil{
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil{
				return first.Next
			}
		}
		next := tail.Next
		cur, tail = Reverse(cur, tail)
		pre.Next = cur
		tail.Next = next
		pre = tail
		cur = tail.Next
	}
	return first.Next
}

func Reverse(head *ListNode, tail *ListNode) (*ListNode, *ListNode) {
	pre := tail.Next
	tmp := head
	for pre != tail{
		next := tmp.Next
		tmp.Next = pre
		pre = tmp
		tmp = next
	}
	return tail, head
}
