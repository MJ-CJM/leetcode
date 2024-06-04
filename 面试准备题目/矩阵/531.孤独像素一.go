// -*- coding:utf-8 -*-
// @Time : 2024/5/5 22:06
// @Author: MJ-CJM
// @File : leetcode/531.孤独像素一
package main

func findLonelyPixel(picture [][]byte) int {
	m := len(picture)
	n := len(picture[0])
	mMap := make([]int, m)
	nMap := make([]int, n)
	res := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if picture[i][j] == 'B' {
				mMap[i]++
				nMap[j]++
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if picture[i][j] == 'B' && mMap[i] == 1 && nMap[j] == 1 {
				res++
			}
		}
	}


	return res
}
