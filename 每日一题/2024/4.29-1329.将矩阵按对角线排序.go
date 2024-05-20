// -*- coding:utf-8 -*-
// @Time : 2024/4/29 23:16
// @Author: MJ-CJM
// @File : leetcode/4.29-1329.将矩阵按对角线排序
package main

import "sort"

func diagonalSort(mat [][]int) [][]int {
	n := len(mat)
	if n <= 0 {
		return mat
	}

	m := len(mat[0])

	digMap := make(map[int][]int)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			key := i - j
			digMap[key] = append(digMap[key], mat[i][j])
		}
	}

	for _, v := range digMap {
		sort.Ints(v)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			key := i - j
			mat[i][j] = digMap[key][0]
			digMap[key] = digMap[key][1:]
		}
	}

	return mat
}


