package main

import "container/heap"

// 分治法
func mergeKLists(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}
	return merge_k(lists, 0, len(lists)-1)
}

func merge_k(lists []*ListNode, start, end int) *ListNode {
	if start == end {
		return lists[start]
	}
	if start > end {
		return nil
	}
	mid := start + (end-start)>>1
	left := merge_k(lists, start, mid)
	right := merge_k(lists, mid+1, end)
	return mergeTwoLists_2(left, right)
}

func mergeTwoLists_2(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var head, node *ListNode
	if l1.Val < l2.Val {
		head = l1
		node = l1
		l1 = l1.Next
	} else {
		head = l2
		node = l2
		l2 = l2.Next
	}
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			node.Next = l1
			l1 = l1.Next
		} else {
			node.Next = l2
			l2 = l2.Next
		}
		node = node.Next
	}
	if l1 != nil {
		node.Next = l1
	}
	if l2 != nil {
		node.Next = l2
	}
	return head
}

// 最小堆法
func mergeKLists_2(lists []*ListNode) *ListNode {
	if lists == nil || len(lists) == 0 {
		return nil
	}
	var h IntHeap
	heap.Init(&h)

	for _, list := range lists {
		if list != nil {
			heap.Push(&h, list)
		}
	}

	dummy := &ListNode{}
	p := dummy

	for h.Len() > 0 {
		min := heap.Pop(&h).(*ListNode)
		p.Next = min
		p = p.Next
		if min.Next != nil {
			heap.Push(&h, min.Next)
		}
	}
	return dummy.Next
}

type IntHeap []*ListNode

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(*ListNode))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
