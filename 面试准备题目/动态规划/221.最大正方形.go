package main

func maximalSquare(matrix [][]byte) int {
	row := len(matrix)
	col := len(matrix[0])
	maxsize := 0

	dp := make([][]int, row + 1)
	for i := 0; i <= row; i++ {
		dp[i] = make([]int, col + 1)
	}

	for i := 1; i <= row; i++ {
		for j := 1; j <= col; j++ {
			if matrix[i-1][j-1] == '1' {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				if dp[i][j] > maxsize {
					maxsize = dp[i][j]
				}
			}
		}
	}
	return maxsize * maxsize
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}
