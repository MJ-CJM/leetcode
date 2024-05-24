// -*- coding:utf-8 -*-
// @Time : 2024/5/8 23:03
// @Author: MJ-CJM
// @File : leetcode/1060.有序数组中缺失的元素
package main

func missingElement(nums []int, k int) int {
	left, right := 0, len(nums)-1

	// 检查整个数组中缺失的数字数量
	totalMissing := nums[right] - nums[0] - right
	if totalMissing < k {
		// 超过数组的范围，直接计算从数组末尾开始的缺失数字
		return nums[right] + k - totalMissing
	}

	// 使用二分查找找到位置
	for left < right {
		mid := left + (right-left)/2

		// 缺失的数量
		missing := nums[mid] - nums[0] - mid

		if missing < k {
			left = mid + 1
		} else {
			right = mid
		}
	}

	// 最终计算第 k 个缺失的数字
	return nums[left-1] + k - (nums[left-1] - nums[0] - (left - 1))
}
