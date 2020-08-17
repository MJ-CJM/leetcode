package main

import "math"

// 动态规划
func profit2(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	mp := make([][][]int, n)
	for i := 0; i < n; i++ {
		mp[i] = make([][]int, 3)
		for j := 0; j < 3; j++ {
			mp[i][j] = make([]int, 2)
		}
	}
	mp[0][0][0], mp[0][0][1] = 0, -prices[0]
	mp[0][1][0], mp[0][1][1] = math.MinInt16, math.MinInt16
	mp[0][2][0], mp[0][2][1] = math.MinInt16, math.MinInt16
	for i := 1; i < n; i++ {
		mp[i][0][0] = mp[i-1][0][0]
		mp[i][0][1] = max(mp[i-1][0][1], mp[i-1][0][0]-prices[i])

		mp[i][1][0] = max(mp[i-1][1][0], mp[i-1][0][1]+prices[i])
		mp[i][1][1] = max(mp[i-1][1][1], mp[i-1][1][0]-prices[i])

		mp[i][2][0] = max(mp[i-1][2][0], mp[i-1][1][1]+prices[i])
	}
	return max_3(mp[n-1][0][0], mp[n-1][1][0], mp[n-1][2][0])
}
