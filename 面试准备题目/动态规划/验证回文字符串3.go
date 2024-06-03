package main

/*
应该是双端dp
dp[i][j][k] 表示s[i : j] 是否是 k 回文
dp[i + 1][j][k] 假如dp[i][j][k]是，dp[i - 1][j]k + 1]肯定是，dp[i][j + 1][k + 1]也是
dp[i][j][k] = dp[i + 1][j][k - 1] || dp[i][j - 1][k - 1]
if s[i] == s[j], dp[i][j][k] = dp[i][j][k] || dp[i + 1][j - 1][k]
*/
func isValidPalindrome(s string, k int) bool {
	dp := make([][][]bool, len(s))
	for i := range dp {
		dp[i] = make([][]bool, len(s))
		for j := range dp[i] {
			dp[i][j] = make([]bool, k + 1)
		}
	}

	for i := len(s) - 1; i >= 0; i -- {
		for j := i; j <= len(s) - 1; j ++ {
			for kk := 0; kk <= min(k, j - i + 1); kk ++ {
				if j == i || s[i] == s[j] && j == i + 1 {
					dp[i][j][kk] = true
				} else {
					if kk > 0 {
						dp[i][j][kk] = dp[i][j][kk] || dp[i + 1][j][kk - 1] || dp[i][j - 1][kk - 1]
					}
					if s[i] == s[j] && i + 1 <= j - 1{
						dp[i][j][kk] = dp[i][j][kk] || dp[i + 1][j - 1][kk]
					}
				}
			}
		}
	}

	return dp[0][len(s) - 1][k]
}
