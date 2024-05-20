// -*- coding:utf-8 -*-
// @Time : 2024/5/16 00:27
// @Author: MJ-CJM
// @File : leetcode/70.爬楼梯
package main

func climbStairs(n int) int {
	if n < 3 {
		return n
	}


	dp := make([]int, n + 1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= n; i++ {
		dp[i] = dp[i - 1] + dp[i - 2]
	}

	return dp[n]
}

// 空间 O 1
func climbStairs2(n int) int {
	if n <= 2{
		return 2
	}
	f1 := 1
	f2 := 2
	f3 := 3
	for i := 3; i <= n; i++ {
		f3 = f1 + f2
		f1 = f2
		f2 = f3
	}
	return f3
}