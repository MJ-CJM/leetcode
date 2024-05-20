// -*- coding:utf-8 -*-
// @Time : 2024/5/16 23:59
// @Author: MJ-CJM
// @File : leetcode/152.乘积最大子数组
package main

// 回溯
func maxProduct2(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	res := nums[0] * nums[1]
	iterm := 1
	level := 0
	maxRecur(level, iterm, n, nums, &res)
	return res
}

func maxRecur(level, iterm, n int, nums []int, res *int) {
	if level == n {
		if iterm > *res {
			*res = iterm
		}
		return
	}
	if iterm > *res && level != 0 {
		*res = iterm
	}


	c1 := iterm * nums[level]
	c2 := nums[level]
	maxRecur(level+1, c1, n, nums, res)
	maxRecur(level+1, c2, n, nums, res)
}


func max_count(x, y int) int {
	if x > y {
		return x
	}
	return y
}



/*
   动态规划：
   1. 定义 dp[i][2], dp[i][0]: 负的最大值，dp[i][1]: 正的最大值
   2. 状态转移公式：if nums[i] >= 0 {
       dp[i][0] = dp[i-1][0] * nums[i],
       dp[i][1] = dp[i-1][1] * nums[i]
   } else {
       dp[i][0] = dp[i-1][1] * nums[i],
       dp[i][1] = dp[i-1][0] * nums[i],
   }
   3. 最后取 dp[n-1][1]
*/
func maxProduct(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][1] = nums[0]
	dp[0][0] = nums[0]
	res := nums[0]

	for i := 1; i < n; i++ {
		if nums[i] >= 0 {
			dp[i][0] = dp[i-1][0] * nums[i]
			dp[i][1] = max(dp[i-1][1] * nums[i], nums[i])
		} else {
			dp[i][0] = min(dp[i-1][1] * nums[i], nums[i])
			dp[i][1] = dp[i-1][0] * nums[i]
		}
		res = max(res, dp[i][1])
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

