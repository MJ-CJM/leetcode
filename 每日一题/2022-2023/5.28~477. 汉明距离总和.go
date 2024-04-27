// -*- coding:utf-8 -*-
// @Time : 2021/5/28 12:24 上午
// @Author: MJ-CJM
// @File : leetcode/5.28~477. 汉明距离总和
package main

import "math/bits"

// 暴力解法
func totalHammingDistance_1(nums []int) int {
	res := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			// res += hammingDistance(nums[i], nums[j])
			// 优化后使用内置函数
			res += bits.OnesCount(uint(nums[i] ^ nums[j]))
		}
	}
	return res
}

// 逐位统计法
func totalHammingDistance(nums []int) int {
	n := len(nums)
	res := 0
	for i := 0; i < 30; i++ {
		c := 0
		for _, val := range nums {
			c += (val >> i) & 1
		}
		res += c * (n - c)
	}
	return res
}
