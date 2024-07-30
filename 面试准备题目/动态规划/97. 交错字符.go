package main

/*
定义数组：定义二维数组 dp，其中 dp[i][j] 表示 s1 的前 i 个字符和 s2 的前 j 个字符能否交错组成 s3 的前 i+j 个字符。

初始化：

dp[0][0] 初始化为 true，表示两个空字符串可以交错组成空字符串。
第一行 dp[i][0] 表示 s2 为空时，只需比较 s1 和 s3 的字符是否相同。
第一列 dp[0][j] 表示 s1 为空时，只需比较 s2 和 s3 的字符是否相同。
状态转移：

对于 dp[i][j]，可以由 dp[i-1][j]（表示 s1 的前 i-1 个字符和 s2 的前 j 个字符能否交错组成 s3 的前 i+j-1 个字符，并且 s1 的第 i 个字符等于 s3 的第 i+j 个字符）或者 dp[i][j-1]（类似地，表示 s1 的前 i 个字符和 s2 的前 j-1 个字符能否交错组成 s3 的前 i+j-1 个字符，并且 s2 的第 j 个字符等于 s3 的第 i+j 个字符）推导而来。
返回结果：最终返回 dp[len1][len2]，即 s1 的所有字符和 s2 的所有字符能否完全交错组成 s3。
*/
func isInterleave(s1 string, s2 string, s3 string) bool {
	len1, len2, len3 := len(s1), len(s2), len(s3)
	if len1 + len2 != len3 {
		return false
	}

	// dp[i][j] 表示 s1 的前 i 个字符和 s2 的前 j 个字符能否交错组成 s3 的前 i+j 个字符
	dp := make([][]bool, len1+1)
	for i := range dp {
		dp[i] = make([]bool, len2+1)
	}

	// 初始条件
	dp[0][0] = true

	// 初始化第一行，s2 为空时
	for i := 1; i <= len1; i++ {
		dp[i][0] = dp[i-1][0] && s1[i-1] == s3[i-1]
	}

	// 初始化第一列，s1 为空时
	for j := 1; j <= len2; j++ {
		dp[0][j] = dp[0][j-1] && s2[j-1] == s3[j-1]
	}

	// 填充 dp 数组
	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			dp[i][j] = (dp[i-1][j] && s1[i-1] == s3[i+j-1]) || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
		}
	}

	return dp[len1][len2]
}

