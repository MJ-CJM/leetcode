package main

// 动态规划，重复子问题
// 状态定义：dp[i][j] j = 0/1 0:没有股票，1：有一股股票
// 状态转移方程：dp[i][0] = max(dp[i-1][0], dp[i-1][1] + a[i] - fee)
//				 dp[i][1] = max(dp[i-1][1], dp[i-1][0] - a[i])
func maxProfit_fee(prices []int, fee int) int {
	n := len(prices)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][1] + prices[i] - fee)
		dp[i][1] = max(dp[i-1][1], dp[i-1][0] - prices[i])
	}
	return dp[n-1][0]
}
