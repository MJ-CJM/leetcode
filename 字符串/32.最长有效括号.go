package main

// 动态规划：重复子问题
// 定义状态：dp[i]，表示以下标 i 字符结尾的最长有效括号的长度
// dp[i] = dp[i-2]+2      s[i]=')' && s[i-1]='('
//         dp[i-1]+d[i-dp[i-1]-2]+2            s[i]=')' && s[i-1]=')' && dp[i-dp[i-1]-1] ='('
func longestValidParentheses(s string) int {
	maxAns := 0
	dp := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if s[i-1] == '(' {
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					dp[i] = 2
				}
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				} else {
					dp[i] = dp[i-1] + 2
				}
			}
		}
		maxAns = max_int(maxAns, dp[i])
	}
	return maxAns
}
