// -*- coding:utf-8 -*-
// @Time : 2024/5/18 11:47
// @Author: MJ-CJM
// @File : leetcode/300.最长递增子序列
package main

/*
   动态规划：
   1. 定义：dp[i] 以 i 为结尾的最长子串
   2. 状态转移方程 dp[i] = nums[i] > nums[j]     max(dp[i], dp[j] + 1)
                            nums[i] <= nums[j]    dp[i]
   3. 结果: max(dp[n-1][j])
*/
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	// dp 初始化
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	res := 0
	// dp 转移方程
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j] + 1)
			}
		}
		if dp[i] > res {
			res = dp[i]
		}
	}

	return res
}


