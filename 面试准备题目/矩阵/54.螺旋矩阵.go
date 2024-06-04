package main

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}

	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1
	result := []int{}

	for top <= bottom && left <= right {
		// 从左到右遍历当前顶部行
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++

		// 从上到下遍历当前右侧列
		for i := top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--

		if top <= bottom {
			// 从右到左遍历当前底部行
			for i := right; i >= left; i-- {
				result = append(result, matrix[bottom][i])
			}
			bottom--
		}

		if left <= right {
			// 从下到上遍历当前左侧列
			for i := bottom; i >= top; i-- {
				result = append(result, matrix[i][left])
			}
			left++
		}
	}

	return result
}
