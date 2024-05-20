// -*- coding:utf-8 -*-
// @Time : 2024/5/19 22:53
// @Author: MJ-CJM
// @File : leetcode/heap_mini
package main

// 建立小顶堆结构体
type MinHeap struct {
	array []int
	size  int
}

// 初始化小顶堆
func NewMinHeap() *MinHeap {
	return &MinHeap{
		array: make([]int, 0),
		size:  0,
	}
}

// 向小顶堆中插入元素
func (h *MinHeap) Insert(val int) {
	h.array = append(h.array, val)
	h.size++
	h.shiftUp(h.size - 1)
}

// 上移操作，使得小顶堆恢复性质
func (h *MinHeap) shiftUp(idx int) {
	for idx > 0 {
		parentIdx := (idx - 1) / 2
		if h.array[idx] >= h.array[parentIdx] {
			break
		}
		h.array[idx], h.array[parentIdx] = h.array[parentIdx], h.array[idx]
		idx = parentIdx
	}
}

// 弹出小顶堆中的最小元素
func (h *MinHeap) Pop() int {
	if h.size == 0 {
		return -1
	}
	minVal := h.array[0]
	h.size--
	h.array[0] = h.array[h.size]
	h.array = h.array[:h.size]
	h.shiftDown(0)
	return minVal
}

// 下移操作，使得小顶堆恢复性质
func (h *MinHeap) shiftDown(idx int) {
	for idx < h.size {
		leftChildIdx := idx*2 + 1
		rightChildIdx := idx*2 + 2
		smallestIdx := idx
		if leftChildIdx < h.size && h.array[leftChildIdx] < h.array[smallestIdx] {
			smallestIdx = leftChildIdx
		}
		if rightChildIdx < h.size && h.array[rightChildIdx] < h.array[smallestIdx] {
			smallestIdx = rightChildIdx
		}
		if smallestIdx == idx {
			break
		}
		h.array[idx], h.array[smallestIdx] = h.array[smallestIdx], h.array[idx]
		idx = smallestIdx
	}
}
