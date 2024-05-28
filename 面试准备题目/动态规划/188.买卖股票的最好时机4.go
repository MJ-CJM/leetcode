// -*- coding:utf-8 -*-
// @Time : 2024/5/17 22:23
// @Author: MJ-CJM
// @File : leetcode/188.买卖股票的最好时机4
package main

/*
	动态规划：
	1. 定义：dp[i][k][j] k 买卖次数，j: 0/1 0: 没有股票，1：有股票
	2. 状态转移方程：dp[i][k][0] = max(dp[i-1][k][0], dp[i-1][k][1] + prices[i] //卖出)
                   dp[i][k][1] = max(dp[i-1][k][1],  dp[i-1][k-1][0] - prices[i] // 买入)
	3. 结果: max dp[n-1][0-k][0]
*/
func maxProfit4(k int, prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}

	// 初始化 dp
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, k+1)
		for j := 0; j < k+1; j++ {
			dp[i][j] = make([]int, 2)
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < k+1; j++ {
			dp[i][j][0] = 0
			dp[i][j][1] = -prices[0]
		}
	}


	// 动态推导
	for i :=1; i < n; i++ {
		for j := 1; j < k+1; j++ {
			dp[i][j][0] = max_count(dp[i-1][j][0], dp[i-1][j][1] + prices[i])
			dp[i][j][1] = max_count(dp[i-1][j][1], dp[i-1][j-1][0] - prices[i])
		}
	}

	// 获取结果
	res := 0
	for j := 0; j < k+1; j++ {
		if dp[n-1][j][0] > res {
			res = dp[n-1][j][0]
		}
	}

	return res
}

func max_count3(x, y int) int {
	if x < y {
		return y
	}
	return x
}