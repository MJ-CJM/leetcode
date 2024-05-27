// -*- coding:utf-8 -*-
// @Time : 2024/5/19 22:51
// @Author: MJ-CJM
// @File : leetcode/1167.链接木棍的最低费用
package main

import "container/heap"

// 建立小顶堆
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

// 计算连接所有木棍的最小成本
func connectSticks(sticks []int) int {
	if len(sticks) == 0 {
		return 0
	}

	// 初始化小顶堆
	h := &MinHeap{}
	heap.Init(h)

	// 将木棍长度数组转换为小顶堆
	for _, stick := range sticks {
		heap.Push(h, stick)
	}

	// 计算总成本
	totalCost := 0
	for h.Len() > 1 {
		// 弹出两个最小长度的木棍
		min1 := heap.Pop(h).(int)
		min2 := heap.Pop(h).(int)
		// 计算连接后的总成本
		cost := min1 + min2
		totalCost += cost
		// 将连接后的木棍长度重新插入小顶堆中
		heap.Push(h, cost)
	}

	return totalCost
}


