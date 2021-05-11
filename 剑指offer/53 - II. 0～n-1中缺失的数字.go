// -*- coding:utf-8 -*-
// @Time : 2021/5/11 11:38 上午
// @Author: MJ-CJM
// @File : leetcode/53 - II. 0～n-1中缺失的数字
package main

func missingNumber(nums []int) int {
	n := len(nums)
	tmp := 0
	if nums[0] != 0 {
		return 0
	}
	for i := 1; i < n; i++ {
		tmp++
		if tmp != nums[i] {
			return tmp
		}
	}
	return n
}
