// -*- coding:utf-8 -*-
// @Time : 2024/5/5 23:10
// @Author: MJ-CJM
// @File : leetcode/163.缺失的区间
package main

func findMissingRanges(nums []int, lower int, upper int) [][]int {
	result := [][]int{}
	// 如果数组为空，则直接返回整个范围
	if len(nums) == 0 {
		return append(result, []int{lower, upper})
	}

	// 检查前端是否有缺失
	if nums[0] > lower {
		result = append(result, []int{lower, nums[0] - 1})
	}

	// 检查数组中的缺失
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] + 1 {
			result = append(result, []int{nums[i-1] + 1, nums[i] - 1})
		}
	}

	// 检查后端是否有缺失
	if nums[len(nums)-1] < upper {
		result = append(result, []int{nums[len(nums)-1] + 1, upper})
	}

	return result

}


