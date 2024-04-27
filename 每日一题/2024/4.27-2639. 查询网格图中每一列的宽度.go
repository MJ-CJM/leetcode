// -*- coding:utf-8 -*-
// @Time : 2024/4/27 21:51
// @Author: MJ-CJM
// @File : leetcode/4.27-2639. 查询网格图中每一列的宽度
package main

import "strconv"

func findColumnWidth(grid [][]int) []int {
	var res []int = []int{}

	m := len(grid)
	if m < 0 {
		return res
	}
	n := len(grid[0])

	for i := 0; i < n; i++ {
		count := 0
		for j := 0; j < m; j++ {
			tmp := len(strconv.Itoa(grid[j][i]))
			if tmp > count {
				count = tmp
			}
		}
		res = append(res, count)
	}
	return res
}
