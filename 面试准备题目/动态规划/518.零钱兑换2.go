// -*- coding:utf-8 -*-
// @Time : 2024/5/19 20:52
// @Author: MJ-CJM
// @File : leetcode/518.零钱兑换2
package main


func change(amount int, coins []int) int {
	dp := make([]int, amount + 1)
	dp[0] = 1 // 初始化总金额为0时有1种组合方法

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}

	return dp[amount]
}

