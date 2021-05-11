// -*- coding:utf-8 -*-
// @Time : 2021/5/11 11:36 上午
// @Author: MJ-CJM
// @File : leetcode/53 - I. 在排序数组中查找数字 I
package main

func search(nums []int, target int) int {
	res := 0
	for _, v := range nums {
		if v == target {
			res++
		}
	}
	return res
}
