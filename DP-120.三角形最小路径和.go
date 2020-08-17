package main

// 暴力递归解法
//func minimumTotal(triangle [][]int) int {
//	result := 0
//	n := len(triangle)
//	i := 0
//	j := 0
//	for i := 0; i < n; i++ {
//		result += triangle[i][0]
//	}
//	iterm := triangle[0][0]
//	_dfsmun(triangle, i, j, n, iterm, &result)
//	return result
//}
//
//func _dfsmun(tri [][]int, i int, j int, n int, iterm int, result *int) {
//	// terminator
//	if i == n-1{
//		if iterm < *result{
//			*result = iterm
//		}
//		return
//	}
//	// process && drill down
//	c1 := iterm + tri[i+1][j]
//	c2 := iterm + tri[i+1][j+1]
//	_dfsmun(tri, i+1, j, n, c1, result)
//	_dfsmun(tri, i+1, j+1, n, c2, result)
//}

// 动态规划
// 重复子问题：problem(i,j) = min(sub(i+1,j), sub(i+1,j+1))+a[i,j]
// 定义状态数组： f[i,j]
// 状态转移方程：f[i, j] = min(f[i+1, j], f[i+1,j+1]) + a[i, j]
func minimumTotal(triangle [][]int) int {
	dp := triangle
	m := len(triangle)
	for i := m - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			dp[i][j] += min_int(dp[i+1][j], dp[i+1][j+1])
		}
	}
	return dp[0][0]
}

func min_int(i int, i2 int) int {
	if i < i2{
		return i
	}else{
		return i2
	}
}
