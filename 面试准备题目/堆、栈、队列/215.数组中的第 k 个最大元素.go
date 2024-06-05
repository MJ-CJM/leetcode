package main

import "container/heap"

type MiniHeap []int

func (h MiniHeap) Len()int {return len(h)}
func (h MiniHeap) Less(i, j int) bool {return h[i] < h[j]}
func (h MiniHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}

func (h *MiniHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MiniHeap) Pop() interface{} {
	n :=  len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

// 主函数，返回第 k 大的元素
func findKthLargest(nums []int, k int) int {
	h := &MiniHeap{}
	heap.Init(h)

	n := len(nums)

	for i := 0; i < k; i++ {
		heap.Push(h, nums[i])
	}

	for i := k; i < n; i++ {
		if nums[i] > (*h)[0] {
			heap.Pop(h)
			heap.Push(h, nums[i])
		}
	}
	res := heap.Pop(h)
	return res.(int)
}
