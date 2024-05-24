// -*- coding:utf-8 -*-
// @Time : 2024/5/23 20:37
// @Author: MJ-CJM
// @File : leetcode/53.最大子数组和
package main

// 动态规划
// 重复子问题：max_sum(i) = Max(max_sum(i-1), 0) + a[i]
// 定义状态数组：f[i]
// 状态转移方程：f[i] = max(f[i-1],0) + a[i]
func maxSubArray(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	result := 0
	for i := 0; i < n; i++ {
		if i == 0{
			dp[i] = nums[i]
			result = dp[i]
		}else{
			dp[i] = max(dp[i-1]+nums[i], nums[i])
			if dp[i] > result{
				result = dp[i]
			}
		}
	}
	return result
}

