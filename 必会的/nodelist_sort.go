package main

// ListNode 定义链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

// sortList 快速排序链表的主函数
func quick_sortList(head *ListNode) *ListNode {
	// 基本情况：如果链表为空或只有一个节点，则不需要排序，直接返回
	if head == nil || head.Next == nil {
		return head
	}

	// 选择基准点
	pivot := head.Val

	// 创建三个链表，分别存储小于、等于和大于基准点的元素
	var less, equal, greater *ListNode
	lessDummy, equalDummy, greaterDummy := &ListNode{}, &ListNode{}, &ListNode{}
	less, equal, greater = lessDummy, equalDummy, greaterDummy


	// 遍历链表进行分区操作
	current := head
	for current != nil {
		if current.Val < pivot {
			less.Next = current
			less = less.Next
		} else if current.Val > pivot {
			greater.Next = current
			greater = greater.Next
		} else {
			equal.Next = current
			equal = equal.Next
		}
		current = current.Next
	}

	// 断开链表尾部的 next 指针
	less.Next, equal.Next, greater.Next = nil, nil, nil

	// 递归排序 less 和 greater 链表
	lessSorted := quick_sortList(lessDummy.Next)
	greaterSorted := quick_sortList(greaterDummy.Next)

	// 合并三部分链表
	return concatenate(lessSorted, equalDummy.Next, greaterSorted)
}

// concatenate 将三个链表合并
func concatenate(less, equal, greater *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	// 合并 less 链表
	current.Next = less
	for current.Next != nil {
		current = current.Next
	}

	// 合并 equal 链表
	current.Next = equal
	for current.Next != nil {
		current = current.Next
	}

	// 合并 greater 链表
	current.Next = greater

	return dummy.Next
}

// 归并排序
// sortList 排序链表的主函数
func sortList(head *ListNode) *ListNode {
	// 基本情况：如果链表为空或只有一个节点，则不需要排序，直接返回
	if head == nil || head.Next == nil {
		return head
	}

	// 使用快慢指针找到链表的中间节点
	slow, fast := head, head
	var prev *ListNode
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 将链表拆分成两部分
	prev.Next = nil

	// 递归排序两部分
	left := sortList(head)
	right := sortList(slow)

	// 合并两部分已排序的链表
	return mergeTwoLists(left, right)
}

// mergeTwoLists 合并两个有序链表
func mergeTwoLists(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}

	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}

	return dummy.Next
}