// -*- coding:utf-8 -*-
// @Time : 2024/5/17 01:45
// @Author: MJ-CJM
// @File : leetcode/121.买卖股票的最好时机
package main

// 常规解法
func maxProfit(prices []int) int {
	res := 0
	m := prices[0]
	for i := 1; i < len(prices); i++ {
		tmp := prices[i] - m
		if tmp > res {
			res = tmp
		}
		if prices[i] < m {
			m = prices[i]
		}
	}
	return res
}

/*
	动态规划：
	1. 定义：dp[i][2]
	2. 状态转移方程：dp[i][1] = prices[i] - dp[i][0]
	3. 结果: min(dp[i][1])
*/
func maxProfit1_2(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}

	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = prices[0]
	dp[0][1] = 0
	res := 0

	for i := 1; i < n; i++ {
		dp[i][0] = min_count(prices[i], dp[i-1][0])
		dp[i][1] = prices[i] - dp[i-1][0]
		if dp[i][1] > res {
			res = dp[i][1]
		}
	}

	return res
}

func min_count2(x, y int) int {
	if x < y {
		return x
	}
	return y
}
