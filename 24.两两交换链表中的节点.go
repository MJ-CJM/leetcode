package main

//type ListNode struct {
//    Val int
//    Next *ListNode
//}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	cur := head
	pre := head.Next
	cur.Next = swapPairs(pre.Next)
	pre.Next = cur
	return pre
}