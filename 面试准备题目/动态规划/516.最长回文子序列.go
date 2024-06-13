package main

func longestPalindromeSubseq(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}

	// 初始化 dp 数组
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1 // 单个字符的回文子序列长度为1
	}

	// 填充 dp 表，从长度为2的子序列开始
	for length := 2; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				dp[i][j] = max(dp[i+1][j], dp[i][j-1])
			}
		}
	}

	// 返回整个字符串的最长回文子序列长度
	return dp[0][n-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
