// -*- coding:utf-8 -*-
// @Time : 2024/5/14 00:26
// @Author: MJ-CJM
// @File : leetcode/994.烂的橘子
package 图

func orangesRotting(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	res := 0
	flag := 0
	queue := [][]int{}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				flag++
				continue
			}
			if grid[i][j] == 2 {
				queue = append(queue, []int{i, j})
			}
		}
	}

	// 如果没有新鲜橘子，直接返回0
	if flag == 0 {
		return 0
	}

	for len(queue) > 0 {
		tmpQueue := [][]int{}

		for _,v := range queue {
			deal := v

			x := deal[0]
			y := deal[1]

			if x - 1 >= 0 && grid[x - 1][y] == 1 {
				grid[x - 1][y] = 2
				flag--
				tmpQueue = append(tmpQueue, []int{x - 1, y})
			}

			if x + 1 < n && grid[x + 1][y] == 1 {
				grid[x + 1][y] = 2
				flag--
				tmpQueue = append(tmpQueue, []int{x + 1, y})
			}

			if y - 1 >= 0 && grid[x][y - 1] == 1 {
				grid[x][y - 1] = 2
				flag--
				tmpQueue = append(tmpQueue, []int{x, y - 1})
			}

			if y + 1 < m && grid[x][y + 1] == 1 {
				grid[x][y + 1] = 2
				flag--
				tmpQueue = append(tmpQueue, []int{x, y + 1})
			}
		}
		if len(tmpQueue) > 0 {
			res++
		}
		queue = tmpQueue
	}

	if flag == 0 {
		return res
	}

	return -1
}