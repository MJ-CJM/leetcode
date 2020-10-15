package main

func spiralOrder(matrix [][]int) []int {
	result := []int{}
	row := len(matrix)
	if row < 1 {
		return result
	}
	col := len(matrix[0])
	if col < 1 {
		return result
	}
	left := 0
	right := col-1
	top := 0
	bottom := row-1
	for left <= right && top <= bottom {
		// 从左到右
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		// 从上往下
		for i := top+1; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		// 从右往左
		if top != bottom {
			for i := right-1; i >= left; i-- {
				result = append(result, matrix[bottom][i])
			}
		}
		// 从下往上
		if left != right {
			for i := bottom-1; i > top; i-- {
				result = append(result, matrix[i][left])
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return result
}