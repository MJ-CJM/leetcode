// -*- coding:utf-8 -*-
// @Time : 2024/5/8 23:03
// @Author: MJ-CJM
// @File : leetcode/1060.有序数组中缺失的元素
package main

/*
	1.	nums[left-1]：
	•	left 是二分查找结束后的索引位置，left-1 表示二分查找过程中最后一个检查过的位置的索引。
	•	nums[left-1] 是在 nums 数组中比 left 索引小的最大值。
	2.	nums[left-1] - nums[0]：
	•	这是从 nums[0] 到 nums[left-1] 范围内的最大可能数字数目。
	3.	(left - 1)：
	•	left - 1 是 nums 数组中从 nums[0] 到 nums[left-1] 的元素个数。
	4.	nums[left-1] - nums[0] - (left - 1)：
	•	这是计算从 nums[0] 到 nums[left-1] 之间的缺失数字的数量。这是因为 nums[left-1] - nums[0] 是该范围内可能存在的数字数量，而减去实际的数字数量 left - 1 就得到了缺失的数字数量。
	5.	k - (nums[left-1] - nums[0] - (left - 1))：
	•	k 是目标缺失数字的序号。这部分计算表示第 k 个缺失的数字在当前 nums[left-1] 后还需要多少个缺失数字。
	6.	nums[left-1] + k - (nums[left-1] - nums[0] - (left - 1))：
	•	这是最终的计算结果。表达式 k - (nums[left-1] - nums[0] - (left - 1)) 计算出了从 nums[left-1] 开始还需要多少个缺失数字。因此，nums[left-1] 加上这个数量就得到了第 k 个缺失的正整数。
*/
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
