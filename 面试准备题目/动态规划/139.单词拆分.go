package main


/*
定义状态：

dp[i] 表示字符串 s 的前 i 个字符（s[0:i]）是否可以由字典中的单词拼接而成。
状态转移方程：

对于每个位置 i，我们检查 s[0:i] 是否可以由字典中的单词组成。
如果存在一个 j（0 <= j < i）使得 dp[j] 为 true，并且 s[j:i] 在字典中，则 dp[i] 为 true。
换句话说，如果 s[0:j] 可以由字典中的单词组成，并且 s[j:i] 在字典中，那么 s[0:i] 也可以由字典中的单词组成。
初始化：

dp[0] 为 true，因为空字符串可以被视为被字典单词拼接而成（不使用任何单词）。
最终结果：

dp[len(s)] 表示字符串 s 是否可以由字典中的单词拼接而成。

 */
func wordBreak(s string, wordDict []string) bool {
	// 将 wordDict 转换为一个集合以便于快速查找
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}

	// 初始化 dp 数组，dp[i] 表示 s 的前 i 个字符是否可以由字典中的单词拼接而成
	dp := make([]bool, len(s)+1)
	dp[0] = true

	// 动态规划求解
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]
}
