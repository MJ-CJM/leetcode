// -*- coding:utf-8 -*-
// @Time : 2023/11/26 23:38
// @Author: MJ-CJM
// @File : leetcode/624.数组列表中的最大距离
package main

import "math"

// 暴力求解
func maxDistance(arrays [][]int) int {
	var res int
	for i := 0; i < len(arrays); i++ {
		for j := i + 1; j < len(arrays); j++ {
			i_l := len(arrays[i])
			j_l := len(arrays[j])
			for k := 0; k < i_l; k++ {
				for p := 0; p < j_l; p++ {
					tmp := int(math.Abs(float64(arrays[i][k] - arrays[j][p])))
					if tmp > res {
						res = tmp
					}
				}
			}
		}
	}
	return res
}

// 优化的地方：先求出每行最大，最小的值，然后依次遍历行直接计算结果
func maxDistance2(arrays [][]int) int {

}
