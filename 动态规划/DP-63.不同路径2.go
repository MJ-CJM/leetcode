package main

// 分治法
//func uniquePathsWithObstacles(obstacleGrid [][]int) int {
//	if obstacleGrid[0][0] == 1{
//		return 0
//	}
//	row := len(obstacleGrid)-1
//	col := len(obstacleGrid[0])-1
//	grid := obstacleGrid
//	i := 0
//	j := 0
//	count := 0
//	_dfscount(grid, row, col, i, j, &count)
//	return count
//}
//
//func _dfscount(grid [][]int, row int, col int, i int, j int, count *int) {
//	// terminator
//	if i == row && j == col{
//		*count++
//		return
//	}
//	// process && drill down
//	if i < row || j < col{
//		if i < row && j < col{
//			if grid[i][j+1] != 1{
//				_dfscount(grid, row, col, i, j+1, count)
//			}
//			if grid[i+1][j] != 1{
//				_dfscount(grid, row, col, i+1, j, count)
//			}
//		}else if i < row && j >= col{
//			if grid[i+1][j] != 1{
//				_dfscount(grid, row, col, i+1, j, count)
//			}
//		}else if i >= row && j < col{
//			if grid[i][j+1] != 1{
//				_dfscount(grid, row, col, i, j+1, count)
//			}
//		}
//	}
//}

// 动态规划,递推
// 定义状态：dp[i][j]
// 状态转移方程：dp[i][j] = dp[i-1][j] + dp[i][j-1] (判断元素是否为1， 为1则dp[i][j]=0)
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1{
		return 0
	}
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	result := make([][]int, m)
	if m == 1 && n == 1 && obstacleGrid[0][0] == 0{
		return 1
	}
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
	}
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] != 1 {
			result[i][0] = 1
		} else {
			break
		}
	}
	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] != 1 {
			result[0][j] = 1
		}else{
			break
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] != 1{
				result[i][j] = result[i-1][j] + result[i][j-1]
			}else{
				result[i][j] = 0
			}
		}
	}
	return result[m-1][n-1]
}
