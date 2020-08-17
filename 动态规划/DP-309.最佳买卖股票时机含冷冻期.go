package main

// 动态规划：重复子问题:第i天的最大值
// 定义状态：mp[i][j], j = 0/1 0:没有股票，1：有一股股票
// 状态转移方程：mp[i][0] = max(mp[i-1][0], mp[i-1][1]+a[i]//卖出)
//             mp[i][1] = max(mp]i-1][1], mp[i-2][0]-a[i]//买入)
// 最后求：max(mp[n-1][0])
func maxProfit5(prices []int) int {
	n := len(prices)
	if n == 0{
		return 0
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 2)
	}
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < n; i++ {
		if i == 1{
			dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
			dp[i][1] = max(dp[i-1][1], dp[i-1][0]-prices[i])
			continue
		}
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-2][0]-prices[i])
	}
	return dp[n-1][0]
}
