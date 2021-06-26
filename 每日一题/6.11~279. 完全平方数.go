// -*- coding:utf-8 -*-
// @Time : 2021/6/11 4:55 下午
// @Author: MJ-CJM
// @File : leetcode/6.11~279. 完全平方数
package main

// 递归推导
func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = i
	}
	for i := 2; i <= n; i++ {
		for j := 0; j*j <= i; j++ {
			dp[i] = minInt1(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}

func minInt1(a, b int) int {
	if a < b {
		return a
	}
	return b
}
