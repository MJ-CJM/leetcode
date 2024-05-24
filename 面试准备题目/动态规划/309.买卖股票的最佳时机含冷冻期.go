// -*- coding:utf-8 -*-
// @Time : 2024/5/17 22:24
// @Author: MJ-CJM
// @File : leetcode/309.买卖股票的最佳时机含冷冻期
package main

/*
   动态规划：
   1. 定义：dp[i][j] j: 0/1 0-无股票，1-有股票
   2. 状态转移方程：dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
                  dp[i][1] = max(dp[i-1][1], dp[i-2][0]-prices[i])
   3. 结果：
*/
func maxProfit5(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}

	if n == 2 {
		if prices[1] > prices[0] {
			return prices[1] - prices[0]
		}
		return 0
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = 0
	dp[0][1] = -prices[0]
	dp[1][0] = max_count(dp[0][0], dp[0][1] + prices[1])
	dp[1][1] = max_count(dp[0][1], dp[0][0] - prices[1])

	for i := 2; i < n; i++ {
		dp[i][0] = max_count(dp[i-1][0], dp[i-1][1] + prices[i])
		dp[i][1] = max_count(dp[i-1][1], dp[i-2][0] - prices[i])
	}

	return dp[n-1][0]
}
