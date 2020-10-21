package main

import (
	"fmt"
)

func movingCount(m int, n int, k int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	return dfs_fanwei(m, n, 0, 0, k, dp)
}

func dfs_fanwei(m int, n int, i int, j int, k int, dp [][]int) int {
	if i < 0 || i >= m || j < 0 || j >= n || dp[i][j] == 1 || (sumPos(i) + sumPos(j)) > k {
		return 0
	}
	dp[i][j] = 1

	sum := 1
	sum += dfs_fanwei(m, n, i+1, j, k, dp)
	sum += dfs_fanwei(m, n, i-1, j, k, dp)
	sum += dfs_fanwei(m, n, i, j+1, k, dp)
	sum += dfs_fanwei(m, n, i, j-1, k, dp)
	return sum
}

func sumPos(i int) int {
	res := 0
	for i > 0 {
		res += i % 10
		i = i / 10
	}
	return res
}



func main() {
	m := 3
	n := 2
	k := 17
	fmt.Println(movingCount(m, n, k))
}
