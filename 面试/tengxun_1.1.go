package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func (h *ListNode) append(i int) {
	for h.Next != nil {
		h = h.Next
	}
	h.Next = &ListNode{Val: i}
}

func main () {
	l1_head := &ListNode{}
	l2_head := &ListNode{}
	n := 0
	_, _ = fmt.Scan(&n)
	for i := 0; i < n; i++ {
		x := 0
		_, _ = fmt.Scan(&x)
		l1_head.append(x)
	}
	m := 0
	_, _ = fmt.Scan(&m)
	for i := 0; i < m; i++ {
		x := 0
		_, _ = fmt.Scan(&x)
		l2_head.append(x)
	}
	out := process_lianbiao2(l1_head, l2_head)
	for i := 0; i < len(out); i++ {
		fmt.Printf("%d ", out[i])
	}
}

func process_lianbiao2(l1_head *ListNode, l2_head *ListNode) []int {
	l1_head = l1_head.Next
	l2_head = l2_head.Next
	res := []int{}
	for l1_head != nil && l2_head != nil {
		if l1_head.Val == l2_head.Val {
			res = append(res, l1_head.Val)
			l1_head = l1_head.Next
			l2_head = l2_head.Next
		}else if l1_head.Val > l2_head.Val {
			l1_head = l1_head.Next
		}else{
			l2_head = l2_head.Next
		}
	}
	return res
}
