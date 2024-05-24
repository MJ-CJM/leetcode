// -*- coding:utf-8 -*-
// @Time : 2024/5/20 00:15
// @Author: MJ-CJM
// @File : leetcode/200.岛屿的数量
package main

// dfs
func numIslands(grid [][]byte) int {
	n := len(grid)
	m := len(grid[0])
	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '0' {
				continue
			}
			count++
			dfsIslands(grid, i, j, n, m)
		}
	}
	return count
}

func dfsIslands(grid [][]byte, i, j, n, m int) {
	if i >= n || i < 0 || j >= m || j < 0 {
		return
	}

	if grid[i][j] == '0' {
		return
	}

	grid[i][j] = '0'
	dfsIslands(grid, i + 1, j, n, m)
	dfsIslands(grid, i - 1, j, n, m)
	dfsIslands(grid, i, j - 1, n, m)
	dfsIslands(grid, i, j + 1, n, m)
}

// 并查集
