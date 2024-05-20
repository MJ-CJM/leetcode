// -*- coding:utf-8 -*-
// @Time : 2024/5/19 22:01
// @Author: MJ-CJM
// @File : leetcode/72. 编辑距离
package main

// DP,
// dp[0...i][0...j]   word1[0:i], word2[0:j]之间的编辑距离
// word[i] == word[j]   dp(i, j) = dp(i-1, j-1) // 分治
// word[i] != word[j]   dp(i, j) = min {dp(i-1, j-1) + 1,
//           dp(i-1, j) + 1,
//           dp(i, j-1) + 1}
func minDistance(word1 string, word2 string) int {
	n := len(word1)
	m := len(word2)

	dp := make([][]int, n + 1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m + 1)
	}

	// 空字符串编辑为空字符串的编辑代价为0
	dp[0][0] = 0

	for i := 1; i <= n; i++ {
		dp[i][0] = i
	}

	for j := 1; j <= m; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min_count(dp[i-1][j-1] + 1, dp[i-1][j] + 1)
				dp[i][j] = min_count(dp[i][j], dp[i][j-1] + 1)
			}
		}
	}

	return dp[n][m]
}

func min_count(x, y int) int {
	if x < y {
		return x
	}
	return y
}