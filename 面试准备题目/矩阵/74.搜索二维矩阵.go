package main

func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix)
	if n == 0 {
		return false
	}
	m := len(matrix[0])
	if m == 0 {
		return false
	}
	if target > matrix[n-1][m-1] || target < matrix[0][0] {
		return false
	}
	min, max := 0, n*m-1
	mid, x, y := 0, 0, 0
	for min <= max {
		mid = min + (max - min)>>1
		x = mid / m
		y = mid % m
		if matrix[x][y] < target {
			min = mid + 1
		}else if matrix[x][y] > target {
			max = mid - 1
		}else{
			return true
		}
	}
	return false
}
