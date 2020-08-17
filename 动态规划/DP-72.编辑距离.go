package main

// DP,
// dp[0...i][0...j]   word1[0:i], word2[0:j]之间的编辑距离
// w1: ....x i
// w2:  ...x j
// dp(i, j) = dp(i-1, j-1) // 分治
// w1: ....x i
// w2:  ...y j
// dp(i, j) = min {dp(i-1, j-1) + 1,
//           dp(i-1, j) + 1,
//           dp(i, j-1) + 1}
func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	var dp = make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, n+1)
	}
	// 空字符串编辑为空字符串的编辑代价为0
	dp[0][0] = 0
	// 初始化
	for i := 1; i <= m; i++ {
		dp[i][0] = i
	}
	for i := 1; i <= n; i++ {
		dp[0][i] = i
	}
	// 递推
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			}else{
				dp[i][j] = min_int(dp[i-1][j-1]+1, dp[i-1][j]+1)
				dp[i][j] = min_int(dp[i][j], dp[i][j-1]+1)
			}
		}
	}
	return dp[m][n]
}