package main

// 动态规划，递推公式
func cuttingRope(n int) int {
	dp := make(map[int]int)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		j, k := 1, i-1
		res := 0
		for j <= k {
			res = max_int(res, max_int(j, dp[j]) * max_int(k, dp[k]))
			j++
			k--
		}
		dp[i] = res
	}
	return dp[n]
}

func max_int(x, y int) int {
	if x > y {
		return x
	}
	return y
}