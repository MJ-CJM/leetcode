package main

import (
	"container/heap"
	"sort"
)

type Node struct {
	value int
	count int
}

func topKFrequent(nums []int, k int) []int {
	list := []Node{}
	numsM := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		numsM[nums[i]]++
	}

	for k, v := range numsM {
		list = append(list, Node{
			value: k,
			count: v,
		})
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].count > list[j].count
	})

	res := []int{}
	for i := 0; i < k; i++ {
		res = append(res, list[i].value)
	}
	return res
}

// 堆的解法
// 定义元素和频率的结构体
type Element struct {
	num   int
	freq  int
}

// 定义最小堆
type MinHeap []Element

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].freq < h[j].freq }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Element))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func topKFrequent2(nums []int, k int) []int {
	// 统计每个元素的频率
	frequency := make(map[int]int)
	for _, num := range nums {
		frequency[num]++
	}

	// 使用最小堆维护频率前 k 高的元素
	h := &MinHeap{}
	heap.Init(h)

	for num, freq := range frequency {
		heap.Push(h, Element{num, freq})
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	// 收集结果
	result := make([]int, 0, k)
	for h.Len() > 0 {
		result = append(result, heap.Pop(h).(Element).num)
	}

	return result
}
