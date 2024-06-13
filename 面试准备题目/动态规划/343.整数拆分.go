package main

/*
   动态规划：dp[i]
   状态转移方程：dp[i] = max(dp[i-x] * x, (i - x) * x)

*/
func integerBreak(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	dp := make([]int, n + 1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		cur := 1
		for j := 1; j < i; j++ {
			tmp := max_count(dp[j] * (i - j), j * (i - j))
			if tmp > cur {
				cur = tmp
			}
		}
		dp[i] = cur
	}
	return dp[n]
}

func max_count(x, y int) int {
	if x > y {
		return x
	}
	return y
}
