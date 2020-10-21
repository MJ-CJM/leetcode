package main

// 动态规划
// 定义状态：dp[i][j]
// 状态转移方程：dp[i][j] = if p[j] == "*" : 1.dp[i-1][j] || dp[i][j-2]      matches(s[i], p[j-1])
//			   							   2. dp[i][j-2]                   otherwise
//     	                 if p[j] != "*" : 1. dp[i-1][j-1]                 matches(s[i],p[j])
// 										  2. false                        otherwise
func isMatch(s string, p string) bool {
	m := len(s)
	n := len(p)
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				if matches(s, p, i, j-1) {
					dp[i][j] = dp[i-1][j] || dp[i][j-2]
				}else{
					dp[i][j] = dp[i][j-2]
				}
			}else{
				if matches(s, p, i, j) {
					dp[i][j] = dp[i-1][j-1]
				}else{
					dp[i][j] = false
				}
			}
		}
	}
	return dp[m][n]
}

func matches(s, p string, i, j int) bool {
	if i == 0 {
		return false
	}
	if p[j-1] == '.' {
		return true
	}
	return s[i-1] == p[j-1]
}