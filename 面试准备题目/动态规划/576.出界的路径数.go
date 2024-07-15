package main

/*
   动态规划：定义 dp[i][j][k] 为从位置 (i, j) 开始，移动 k 步出界的路径数
   dp[i][j][k] = dp[i-1][j][k-1] + dp[i+1][j][k-1] + dp[i][j-1][k-1] + dp[i][j+1][k-1]
*/
func findPaths(m int, n int, maxMove int, startRow int, startCol int) int {
	const MOD = 1000000007
	dp := make([][][]int, m+2)
	for i := range dp {
		dp[i] = make([][]int, n+2)
		for j := range dp[i] {
			dp[i][j] = make([]int, maxMove+1)
		}
	}

	for moves := 1; moves <= maxMove; moves++ {
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if i == 0 {
					dp[i][j][moves]++
				}
				if i == m-1 {
					dp[i][j][moves]++
				}
				if j == 0 {
					dp[i][j][moves]++
				}
				if j == n-1 {
					dp[i][j][moves]++
				}
				if i > 0 {
					dp[i][j][moves] = (dp[i][j][moves] + dp[i-1][j][moves-1]) % MOD
				}
				if i < m-1 {
					dp[i][j][moves] = (dp[i][j][moves] + dp[i+1][j][moves-1]) % MOD
				}
				if j > 0 {
					dp[i][j][moves] = (dp[i][j][moves] + dp[i][j-1][moves-1]) % MOD
				}
				if j < n-1 {
					dp[i][j][moves] = (dp[i][j][moves] + dp[i][j+1][moves-1]) % MOD
				}
			}
		}
	}
	return dp[startRow][startCol][maxMove]
}

