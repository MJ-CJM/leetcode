// -*- coding:utf-8 -*-
// @Time : 2024/5/23 00:57
// @Author: MJ-CJM
// @File : leetcode/407.接雨水2
package main

import "container/heap"

/*
思路总结
使用最小堆：

通过最小堆，我们可以每次处理高度最小的单元格，从而确保水从低到高地逐步处理。
边界初始化：

初始化时将所有边界单元格加入堆中，因为边界单元格无法储水，只能用于限制水的流动。
访问和标记：

使用 visited 数组来标记已经访问过的单元格，避免重复处理。
计算接水量：

每次从堆中弹出一个单元格时，检查其四个方向上的邻接单元格，计算该单元格可能接住的水量，然后将未访问的邻接单元格加入堆中。
整体逻辑：

通过以上步骤，利用优先级队列和动态标记访问状态，逐步扩大水池区域并计算总的接水量。
这种方法有效地解决了二维高度图接水的问题，通过优先级队列确保以最低高度开始处理，可以高效地计算出整个图中的最大接水量。
 */


// 定义一个结构体来存储矩阵中的单元格
type Cell struct {
	height, row, col int
}

// 定义一个最小堆
type MinHeap []Cell

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].height < h[j].height }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Cell))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func trapRainWater(heightMap [][]int) int {
	if len(heightMap) == 0 || len(heightMap[0]) == 0 {
		return 0
	}

	m, n := len(heightMap), len(heightMap[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	h := &MinHeap{}
	heap.Init(h)

	// 将边界单元格加入堆中并标记为已访问
	for i := 0; i < m; i++ {
		heap.Push(h, Cell{heightMap[i][0], i, 0})
		heap.Push(h, Cell{heightMap[i][n-1], i, n-1})
		visited[i][0] = true
		visited[i][n-1] = true
	}
	for j := 1; j < n-1; j++ {
		heap.Push(h, Cell{heightMap[0][j], 0, j})
		heap.Push(h, Cell{heightMap[m-1][j], m-1, j})
		visited[0][j] = true
		visited[m-1][j] = true
	}

	directions := []struct{ x, y int }{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	water := 0

	// 开始处理堆中的单元格
	for h.Len() > 0 {
		cell := heap.Pop(h).(Cell)
		for _, d := range directions {
			r, c := cell.row+d.x, cell.col+d.y
			if r >= 0 && r < m && c >= 0 && c < n && !visited[r][c] {
				visited[r][c] = true
				water += max(0, cell.height-heightMap[r][c])
				heap.Push(h, Cell{max(cell.height, heightMap[r][c]), r, c})
			}
		}
	}

	return water
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
解释
定义结构体 Cell：用来存储矩阵中的单元格，包含高度、行和列的信息。

定义最小堆 MinHeap：用于优先级队列操作。

初始化最小堆和访问标记：将边界单元格加入堆中并标记为已访问。

处理堆中的单元格：

从堆中取出高度最低的单元格。
遍历其四个方向上的相邻单元格。
如果相邻单元格没有被访问过，则计算其可以积水的体积，并将其加入堆中。
更新访问标记。
返回最终的积水量：将所有可以积水的体积累加，返回结果。

这个算法通过优先级队列和堆来处理积水问题，确保从最低的地方开始填充水，避免了边界问题和累积问题，保证了计算的准确性。
 */