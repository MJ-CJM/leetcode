// -*- coding:utf-8 -*-
// @Time : 2024/5/24 01:18
// @Author: MJ-CJM
// @File : leetcode/560.和为 k 的子数组
package main

/*
给你一个整数数组 nums 和一个整数 k ，请你统计并返回 该数组中和为 k 的子数组的个数 。

子数组是数组中元素的连续非空序列。
 */
// 解法：后面的前缀 - 前面的前缀
func subarraySum(nums []int, k int) int {
	preSum := 0
	count := 0
	nMap := make(map[int]int)
	nMap[0] = 1
	for _, num := range nums {
		preSum += num
		if v, ok := nMap[preSum - k]; ok {
			count += v
		}
		nMap[preSum]++
	}
	return count
}


