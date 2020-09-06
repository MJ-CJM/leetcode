package main

// 动态规划：重复子问题
// 定义状态：dp[i][j]  i 种面额 j 元
// dp[i][j] = 不选 dp[i-1][j]
// 			   选 dp[i-1][j] + dp[i][j-coins[i-1]]
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
			}else{
				dp[i][j] = dp[i-1][j] + dp[i][j - coin]
			}
		}
	}
	return dp[n][amount]
}

// 一维动态规划
func change_1(amount int, coins []int) int {
	n := len(coins)
	dp := make([]int, amount+1)
	dp[0] = 1
	for i := 0; i < n; i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] = dp[j - coins[i]] + dp[j]
		}
	}
	return dp[amount]
}