// -*- coding:utf-8 -*-
// @Time : 2021/6/10 7:24 下午
// @Author: MJ-CJM
// @File : leetcode/6.10~518. 零钱兑换 II
package main

// 重复子问题
// 定义状态： dp[i][j]    有 i 中面额，总共 j 元。
// dp[i][j] = 不选  dp[i-1][j]
//             选   dp[i-1][j] + dp[i][j-conis[i-1]]

func change(amount int, coins []int) int {
	n := len(coins)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, amount+1)
	}
	for i := 0; i <= n; i++ {
		dp[i][0] = 1
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= amount; j++ {
			coin := coins[i-1]
			if j < coin {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-coin]
			}
		}
	}
	return dp[n][amount]
}
