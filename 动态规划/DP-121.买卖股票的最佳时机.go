package main

func maxProfit(prices []int) int {
	n := len(prices)
	result := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if prices[j]-prices[i] > result {
				result = prices[j] - prices[i]
			}
		}
	}
	return result
}

// 动态规划解法
func profit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	mp := make([][]int, n)
	for i := 0; i < n; i++ {
		mp[i] = make([]int, 3)
	}
	mp[0][0] = 0
	mp[0][1] = -prices[0]
	mp[0][2] = 0
	result := 0
	for i := 1; i < n; i++ {
		mp[i][0] = mp[i-1][0]
		mp[i][1] = max(mp[i-1][1], mp[i-1][0]-prices[i])
		mp[i][2] = mp[i-1][1] + prices[i]
		result = max_4(result, mp[i][0], mp[i][1], mp[i][2])
	}
	return result
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func max_4(x, y, z, p int) int {
	tmp := x
	if y > tmp {
		tmp = y
	}
	if z > tmp {
		tmp = z
	}
	if p > tmp {
		tmp = p
	}
	return tmp
}
