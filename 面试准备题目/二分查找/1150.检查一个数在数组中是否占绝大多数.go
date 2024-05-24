// -*- coding:utf-8 -*-
// @Time : 2024/5/8 23:33
// @Author: MJ-CJM
// @File : leetcode/1150.检查一个数在数组中是否占绝大多数
package main

func isMajorityElement(nums []int, target int) bool {
	left := 0
	right := len(nums) - 1
	mid := 0
	first, second := -1, -1

	// 查找第一次出现的地方
	for left <= right {
		mid = left + (right - left) / 2

		if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if left < len(nums) && nums[left] == target {
		first = left
	}

	// 如果没有找到目标值，直接返回 false
	if first == -1 {
		return false
	}

	left = 0
	right = len(nums) - 1

	// 查找最后一次出现的地方
	for left <= right {
		mid = left + (right - left) / 2

		if nums[mid] > target {
			right = mid - 1
		} else  {
			left = mid + 1
		}
	}
	if right >= 0 && nums[right] == target {
		second = right
	}

	count := second - first + 1
	return count > (len(nums)/2)
}
