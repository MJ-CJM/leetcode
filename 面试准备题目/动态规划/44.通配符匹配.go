package main

import "fmt"

// 动态规划：重复子问题
// 定义状态：dp[i][j]  表示 s 的前 i 个字符与 p 中的前 j 个字符是否能够匹配
// dp[i][j] = p[j] == "*" 1.dp[i][j-1] || dp[i-1][j]
//			  p[j] != "*" 2.dp[i-1][j-1]              s[i] == p[j] || p[j] == '?'
//                        2.false                      otherwise
func isMatch_2(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++{
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 1; i <= n; i++ {
		if p[i-1] == '*' {
			dp[0][i] = true
		}else{
			break
		}
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '*' {
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			}else{
				if p[j-1] == '?' || p[j-1] == s[i-1] {
					dp[i][j] = dp[i-1][j-1]
				}else{
					dp[i][j] = false
				}
			}
		}
	}
	return dp[m][n]
}

func main() {
	s := "sbc"
	a := s+","
	fmt.Println(a)
}
