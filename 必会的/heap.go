// -*- coding:utf-8 -*-
// @Time : 2024/5/19 22:48
// @Author: MJ-CJM
// @File : leetcode/heap
package main

import (
	"container/heap"
	"fmt"
)

package main

import (
"container/heap"
"fmt"
)

// 定义一个大顶堆结构体
type MaxHeap []int

// 实现 heap.Interface 接口的方法
func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push 方法将元素推入堆中
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop 方法从堆中弹出最大元素
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 定义一个小顶堆结构体
type MinHeap []int

// 实现 heap.Interface 接口的方法
func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push 方法将元素推入堆中
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop 方法从堆中弹出最小元素
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	// 示例：建立一个大顶堆
	maxHeap := &MaxHeap{3, 5, 1, 8, 10, 2, 7}
	heap.Init(maxHeap)
	fmt.Printf("Max Heap: %v\n", maxHeap)

	// 示例：建立一个小顶堆
	minHeap := &MinHeap{3, 5, 1, 8, 10, 2, 7}
	heap.Init(minHeap)
	fmt.Printf("Min Heap: %v\n", minHeap)
}

