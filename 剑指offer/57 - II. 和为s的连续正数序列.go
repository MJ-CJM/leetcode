// -*- coding:utf-8 -*-
// @Time : 2021/5/11 7:42 下午
// @Author: MJ-CJM
// @File : leetcode/57 - II. 和为s的连续正数序列
package main

// 暴力求解 n`2
func findContinuousSequence(target int) [][]int {
	res := make([][]int, 0)
	for i := 1; i < target; i++ {
		for j := i; j < target; j++ {
			if (i+j)*(j-i+1)/2 == target {
				tmp := make([]int, 0)
				for p := i; p <= j; p++ {
					tmp = append(tmp, p)
				}
				res = append(res, tmp)
				break
			} else if (i+j)*(j-i+1)/2 > target {
				break
			}
		}
	}
	return res
}

// 滑动窗口的方法 n
func findContinuousSequence_2(target int) [][]int {
	i, j, sum := 1, 1, 0
	res := make([][]int, 0)
	for i <= target>>1 {
		if sum > target {
			sum -= i
			i++
		} else if sum < target {
			sum += j
			j++
		} else {
			tmp := make([]int, 0)
			for p := i; p < j; p++ {
				tmp = append(tmp, p)
			}
			res = append(res, tmp)
			sum -= i
			i++
		}
	}
	return res
}
