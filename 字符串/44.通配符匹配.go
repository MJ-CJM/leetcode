package main
// 动态规划：重复子问题
// 定义状态：dp[i][j]  表示 s 的前 i 个字符与 p 中的前 j 个字符是否能够匹配
// dp[i][j] = p[j] != "*" 1.dp[i-1][j-1]  matches(s[i],p[j])
//		                  1.false	      otherwise
//			  p[j] == "*" 2. dp[i-1][j] or dp[i][j-2]             matches(s[i],p[j-1])
//                        2. dp[i][j-2]                           otherwise
func isMatch_2(s string, p string) bool {

}
