package main

// 动态规划
// 定义状态：dp[i][j]
// 状态转移方程：dp[i][j] = min(dp[i-1][j], dp[i][j-1])
func minPathSum(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
	}
	dp[0][0] = grid[0][0]
	for i := 1; i < n; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for i := 1; i < m; i++ {
		dp[0][i] = dp[0][i-1] + grid[0][i]
	}
	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			dp[i][j] = min_int(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[n-1][m-1]
}
