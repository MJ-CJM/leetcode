// -*- coding:utf-8 -*-
// @Time : 2024/5/19 22:55
// @Author: MJ-CJM
// @File : leetcode/heap_max
package main

// 建立大顶堆结构体
type MaxHeap struct {
	array []int
	size  int
}

// 初始化大顶堆
func NewMaxHeap() *MaxHeap {
	return &MaxHeap{
		array: make([]int, 0),
		size:  0,
	}
}

// 向大顶堆中插入元素
func (h *MaxHeap) Insert(val int) {
	h.array = append(h.array, val)
	h.size++
	h.shiftUp(h.size - 1)
}

// 上移操作，使得大顶堆恢复性质
func (h *MaxHeap) shiftUp(idx int) {
	for idx > 0 {
		parentIdx := (idx - 1) / 2
		if h.array[idx] <= h.array[parentIdx] {
			break
		}
		h.array[idx], h.array[parentIdx] = h.array[parentIdx], h.array[idx]
		idx = parentIdx
	}
}

// 弹出大顶堆中的最大元素
func (h *MaxHeap) Pop() int {
	if h.size == 0 {
		return -1
	}
	maxVal := h.array[0]
	h.size--
	h.array[0] = h.array[h.size]
	h.array = h.array[:h.size]
	h.shiftDown(0)
	return maxVal
}

// 下移操作，使得大顶堆恢复性质
func (h *MaxHeap) shiftDown(idx int) {
	for idx < h.size {
		leftChildIdx := idx*2 + 1
		rightChildIdx := idx*2 + 2
		largestIdx := idx
		if leftChildIdx < h.size && h.array[leftChildIdx] > h.array[largestIdx] {
			largestIdx = leftChildIdx
		}
		if rightChildIdx < h.size && h.array[rightChildIdx] > h.array[largestIdx] {
			largestIdx = rightChildIdx
		}
		if largestIdx == idx {
			break
		}
		h.array[idx], h.array[largestIdx] = h.array[largestIdx], h.array[idx]
		idx = largestIdx
	}
}

