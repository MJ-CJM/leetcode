package main

func printMatrix(matrix [][]int) []int {
	result := []int{}
	row := len(matrix)
	if row == 0 {
		return result
	}
	column := len(matrix[0])
	if column == 0 {
		return result
	}
	//定义下标变量
	top := 0
	bottom := row - 1
	left := 0
	right := column - 1
	for left <= right && top <= bottom {
		//从左往右，行数为top，列数从left开始，+1直到right
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		//从上往下，行数从top+1开始，+1直到bottom，列数为right
		for i := top + 1; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		//row不重复，从右往左，行数为bottom，列数从right-1开始，-1直到left
		if top != bottom {
			for i := right - 1; i >= left; i-- {
				result = append(result, matrix[bottom][i])
			}
		}
		//column不重复，从下往上，行数从bottom-1开始，-1直到top-1，所以没有等号
		if left != right {
			for i := bottom - 1; i > top; i-- {
				result = append(result, matrix[i][left])
			}
		}
		//一圈结束
		left++
		right--
		top++
		bottom--
	}
	return result
}
