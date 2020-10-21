package main

import "fmt"

// 分治法
//func uniquePaths(m int, n int) int {
//	i := 1
//	j := 1
//	count := 0
//	_dfspath(m, n, i, j, &count)
//	return count
//}
//
//func _dfspath(m int, n int, i int, j int, count *int) {
//	if i == m && j == n{
//		*count++
//		return
//	}
//	if i < m && j < n{
//		_dfspath(m, n, i+1, j, count)
//		_dfspath(m, n, i, j+1, count)
//	}
//	if i < m && j >= n{
//		_dfspath(m, n, i+1, j, count)
//	}
//	if i >= m && j < n{
//		_dfspath(m, n, i, j+1, count)
//	}
//}

// 动态规划
func uniquePaths(m int, n int) int {
	if m == 1 && n == 1{
		return 1
	}
	result := make([][]int, m)
	for i := 0; i < m; i ++{
		result[i] = make([]int, n)
	}
	return _dfspath(m-1, n-1, result)
}

func _dfspath(m int, n int, result [][]int) int{
	if (m == 0 && n == 1) || (m == 1 && n == 0){
		result[m][n] = 1
	}
	if result[m][n] == 0{
		if m > 0 && n > 0{
			result[m][n] = _dfspath(m-1, n, result) + _dfspath(m, n-1, result)
		}
		if m == 0 && n > 0{
			result[m][n] = _dfspath(m, n-1, result)
		}
		if m > 0 && n == 0{
			result[m][n] = _dfspath(m-1, n, result)
		}

	}
	return result[m][n]
}

// 递推
// 动态规划
// 定义状态：dp[i][j]
// 状态转移方程：dp[i][j] = dp[i-1][j] + dp[i][j-1]
func uniquePath(m int, n int) int{
	if m <= 0 || n <= 0{
		return 0
	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

func main(){
	m := 2
	n := 3
	fmt.Println(uniquePaths(m, n))
}