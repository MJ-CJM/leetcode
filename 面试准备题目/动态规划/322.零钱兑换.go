// -*- coding:utf-8 -*-
// @Time : 2024/5/19 20:51
// @Author: MJ-CJM
// @File : leetcode/322.零钱兑换
package main

// 递归回溯法
func coinChange2(coins []int, amount int) int {
	iterm := 0
	tmp := 0
	res := amount + 1
	coin(tmp, iterm, &res, coins, amount)
	if res == amount + 1 {
		return -1
	}
	return res
}

func coin(tmp, iterm int, res *int, coins []int, amount int) {
	if tmp >= amount {
		if tmp == amount && iterm < *res {
			*res = iterm
		}
		return
	}

	level := iterm + 1
	for i := 0; i < len(coins); i++ {
		coin(tmp + coins[i], level, res, coins, amount)
	}
}


// 动态规划
// 重复子问题：min(i) = min{min(n-k), for k in range coins} + 1
// 定义状态：f[i]
// 状态转移方程：f[n] = min{f[n-k], for k in range coins} + 1
func coinChange(coins []int, amount int) int {
	// 初始化 dp 数组，dp[i] 表示凑成金额 i 所需的最少硬币数量
	dp := make([]int, amount+1)

	// 动态规划
	for i := 1; i <= amount; i++ {
		dp[i] = -1
		for _, coin := range coins {
			if i < coin || dp[i - coin] == -1 {
				continue
			}
			v := dp[i-coin] + 1
			if dp[i] == -1 || dp[i] > v {
				dp[i] = v
			}
		}
	}


	return dp[amount]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}



