// -*- coding:utf-8 -*-
// @Time : 2024/5/16 00:35
// @Author: MJ-CJM
// @File : leetcode/120.三角形最小路径求和
package main

// 递归回溯
func minimumTotal2(triangle [][]int) int {
	result := 0
	n := len(triangle)
	k := 0
	j := 0
	for i := 0; i < n; i++ {
		result += triangle[i][0]
	}
	iterm := triangle[0][0]
	_dfsmun(triangle, k, j, n, iterm, &result)
	return result
}

func _dfsmun(tri [][]int, i int, j int, n int, iterm int, result *int) {
	// terminator
	if i == n-1{
		if iterm < *result{
			*result = iterm
		}
		return
	}
	// process && drill down
	c1 := iterm + tri[i+1][j]
	c2 := iterm + tri[i+1][j+1]
	_dfsmun(tri, i+1, j, n, c1, result)
	_dfsmun(tri, i+1, j+1, n, c2, result)
}

// 动态规划
/*
   动态规划：
   定义：dp[i][j],从下往上开始遍历
   状态转移方程：dp[i][j] = min(dp[i+1][j], dp[i+1][j+1]) + tri[i][j]
   时间复杂度：O(n*m) 同空间复杂度
*/

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	res := triangle[0][0]
	if n == 1 {
		for i := 0; i < len(triangle[0]); i++ {
			if triangle[0][i] < res {
				res = triangle[0][i]
			}
		}
		return res
	}


	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, len(triangle[i]))
	}
	for j := 0; j < len(triangle[n-1]); j++ {
		dp[n-1][j] = triangle[n-1][j]
	}
	for i := n-2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			dp[i][j] = min_count(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
		}
	}
	return dp[0][0]
}

func min_count(x, y int) int {
	if x > y {
		return y
	}
	return x
}
