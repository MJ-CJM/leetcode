// -*- coding:utf-8 -*-
// @Time : 2024/5/24 00:48
// @Author: MJ-CJM
// @File : leetcode/128.最长连续序列
package main

import "sort"

/*
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
 */
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	nMap := make(map[int]int)
	res := 1
	for _, v := range nums {
		nMap[v] = nMap[v-1] + 1
		if nMap[v] > res {
			res = nMap[v]
		}
	}
	return res
}