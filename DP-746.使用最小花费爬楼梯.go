package main

// 动态规划：重复子问题
// 定义状态
//  状态转移方程：dp[i] = min(dp[i-1], dp[i-2])
func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	cost = append(cost, 0)
	dp := make([]int, n+1)
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i <= n; i++ {
		dp[i] = min_int(dp[i-1], dp[i-2]) + cost[i]
	}
	return dp[n]
}