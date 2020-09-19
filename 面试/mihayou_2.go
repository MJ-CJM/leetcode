package main

import "fmt"

func main() {
	var s string
	var p string
	_, _ = fmt.Scan(&s)
	_, _ = fmt.Scan(&p)
	if isMatch(s, p) {
		fmt.Printf("%s", "true")
	}else{
		fmt.Printf("%s", "false")
	}
}

// 动态规划：重复子问题
// 定义状态：dp[i][j]  表示 s 的前 i 个字符与 p 中的前 j 个字符是否能够匹配
// dp[i][j] = p[j] != "*" 1.dp[i-1][j-1]  matches(s[i],p[j])
//		                  1.false	      otherwise
//			  p[j] == "*" 2. dp[i-1][j] or dp[i][j-2]             matches(s[i],p[j-1])
//                        2. dp[i][j-2]                           otherwise
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '+' {
				if matches(s, p, i, j-1) {
					dp[i][j] = dp[i-1][j] || dp[i][j-2]
				}else{
					dp[i][j] = dp[i][j-2]
				}
				continue
			}
			if p[j-1] == '*' {
				if matches(s, p, i, j-1) {
					dp[i][j] = dp[i-1][j] || dp[i][j-2]
				}else{
					dp[i][j] = dp[i][j-2]
				}
			}else {
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

func matches(s, p string,  i, j int) bool {
	if i == 0 {
		return false
	}
	if p[j-1] == '.' {
		return true
	}
	return s[i-1] == p[j-1]
}

//package main
//
//import "fmt"
//
//func main() {
//	fmt.Printf("%s", "true")
//}