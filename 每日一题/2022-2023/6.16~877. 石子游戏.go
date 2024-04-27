// -*- coding:utf-8 -*-
// @Time : 2021/6/16 12:01 上午
// @Author: MJ-CJM
// @File : leetcode/6.16~877. 石子游戏
package main

// 动态规划
// dp[i][j]：表示剩下的石子堆为下标 i 到 j 时，当前玩家与另外一个玩家石子数量之差的最大值，当前玩家不一定是先手的人。
// 只有当 i <= j 时，剩下的石子堆才有意义，因此当 i > j 时，dp[i][j] = 0
// 当 i = j 时，dp[i][j] = piles[i]
// 当 i < j 时，dp[i][j] = max(piles[i]-dp[i+1][j], piles[j]-dp[i][j-1])
func stoneGame(piles []int) bool {
	n := len(piles)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
		dp[i][i] = piles[i]
	}
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[i][j] = max2(piles[i]-dp[i+1][j], piles[j]-dp[i][j-1])
		}
	}
	return dp[0][n-1] > 0
}

func max2(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
