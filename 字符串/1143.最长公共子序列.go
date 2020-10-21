package main

// 动态规划 重复子问题
// 定义状态： dp[i][j]
// 状态转移方程：若匹配则dp[i][j] = dp[i-1][j-1]+1 若不匹配则dp[i][j] = max(dp[i][j-1],dp[i-1][j])
func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1) + 1)
	for i := 0; i <= len(text1); i++ {
		dp[i] = make([]int, len(text2)+1)
	}
	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			}else{
				dp[i][j] = max_int(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}

func GetCommon( str1 string ,  str2 string ) string {
	dp := make([][]string, len(str1) + 1)
	for i := 0; i <= len(str1); i++ {
		dp[i] = make([]string, len(str2) + 1)
	}
	for i := 1; i <= len(str1); i++ {
		for j := 1; j <= len(str2); j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + string(str1[i-1])
			}else{
				l1 := dp[i-1][j]
				l2 := dp[i][j-1]
				if len(l1) > len(l2) {
					dp[i][j] = dp[i-1][j]
				}else{
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}
	return dp[len(str1)][len(str2)]
}


func max_int(i int, j int) int {
	if i > j {
		return i
	}
	return j
}
