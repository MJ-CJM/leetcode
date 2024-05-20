// -*- coding:utf-8 -*-
// @Time : 2024/5/17 22:24
// @Author: MJ-CJM
// @File : leetcode/714.买卖股票的最佳时机含手续费
package main

/*
   动态规划：
   1. 定义：dp[i][j] j: 0/1 0-无股票，1-有股票
   2. 状态转移方程：dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
                  dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i] - fee)
   3. 结果：dp[n-1][0]
*/
func maxProfit6(prices []int, fee int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}

	dp[0][0] = 0
	dp[0][1] = -(prices[0] + fee)

	for i := 1; i < n; i++ {
		dp[i][0] = max_count(dp[i-1][0], dp[i-1][1] + prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i] - fee)
	}
	return dp[n-1][0]
}

