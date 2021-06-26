// -*- coding:utf-8 -*-
// @Time : 2021/5/15 9:52 上午
// @Author: MJ-CJM
// @File : leetcode/56 - I. 数组中数字出现的次数
package main

// 时间：n，空间：n
func singleNumbers(nums []int) []int {
	res := make([]int, 0)
	res = append(res, nums[0])
	res = append(res, nums[1])
	for i := 2; i < len(nums); i++ {
		flag, index := numIn(res, nums[i])
		if flag {
			res = append(res[:index], res[index+1: len(res)]...)
		} else {
			res = append(res, nums[i])
		}
	}
	return res
}

func numIn(nums []int, target int) (bool, int) {
	tmp := -1
	for k, v := range nums {
		if v == target {
			tmp = k
			return true, tmp
		}
	}
	return false, tmp
}