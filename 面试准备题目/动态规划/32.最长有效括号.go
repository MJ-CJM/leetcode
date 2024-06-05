package main

// 动态规划：重复子问题
// 定义状态：dp[i]，表示以下标 i 字符结尾的最长有效括号的长度
// dp[i] = dp[i-2]+2      s[i]=')' && s[i-1]='('
//         dp[i-1]+d[i-dp[i-1]-2]+2            s[i]=')' && s[i-1]=')' && dp[i-dp[i-1]-1] ='('
func longestValidParentheses(s string) int {
	maxAns := 0  // 初始化最大长度为0
	dp := make([]int, len(s))  // 创建一个与字符串长度相同的dp数组

	for i := 1; i < len(s); i++ {  // 从索引1开始遍历字符串
		if s[i] == ')' {  // 只处理当前字符是 ')' 的情况
			if s[i-1] == '(' {  // 检查前一个字符是否是 '('
				if i >= 2 {
					dp[i] = dp[i-2] + 2  // 匹配的有效括号长度加2
				} else {
					dp[i] = 2  // 如果前面没有更多字符，则当前有效括号长度为2
				}
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {  // 检查前面匹配的有效括号的前一个字符是否是 '('
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2  // 加上前面匹配的有效括号长度
				} else {
					dp[i] = dp[i-1] + 2  // 如果前面没有更多字符，则当前有效括号长度为前面匹配的有效括号长度加2
				}
			}
		}
		maxAns = max_int(maxAns, dp[i])  // 更新最大长度
	}
	return maxAns
}
