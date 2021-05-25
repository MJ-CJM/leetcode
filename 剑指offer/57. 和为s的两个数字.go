// -*- coding:utf-8 -*-
// @Time : 2021/5/11 7:35 下午
// @Author: MJ-CJM
// @File : leetcode/57. 和为s的两个数字
package main

func twoSum(nums []int, target int) []int {
	map_twoSum := make(map[int]bool)
	for _, v := range nums {
		if map_twoSum[target - v] == true {
			return []int{v, target-v}
		} else {
			map_twoSum[v] = true
		}
	}
	return nil
}
