package main

func findNumberIn2DArray(matrix [][]int, target int) bool {
	n := len(matrix)
	i := n-1
	j := 0
	for i > -1 {
		if j < len(matrix[i]) {
			if matrix[i][j] > target {
				i--
			}else if matrix[i][j] < target {
				j++
			}else{
				return true
			}
		}else{
			return false
		}
	}
	return false
}
