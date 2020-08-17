package main

import "fmt"

// 动态规划
// 重复子问题：min(i) = min{min(n-k), for k in range coins} + 1
// 定义状态：f[i]
// 状态转移方程：f[n] = min{f[n-k], for k in range coins} + 1
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = -1
		for  _, v := range coins{
			if i < v || dp[i-v] == -1{
				continue
			}
			count := dp[i-v] + 1
			if dp[i] == -1 || dp[i] > count{
				dp[i] = count
			}
		}
	}
	return dp[amount]
}


// BFS+递归
func coinChange_2(coins []int, amount int) int {
	level := 0
	iterm := 0
	result := 0
	_bfscoin(coins, iterm, level, &result, amount)
	if result != 0{
		return result
	}
	return -1
}

func _bfscoin(coins []int, iterm int, level int, result *int, amount int) {
	// terminator
	if iterm > amount{
		return
	}

	if iterm == amount {
		*result = level
		return
	}
	level++
	for i := 0; i < len(coins); i++ {
		c1 := iterm + coins[i]
		_bfscoin(coins, c1, level, result, amount)
	}
}

func main() {
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println(coinChange(coins, amount))
}
