package main

/*
 * clockwise rotate 顺时针旋转
 * first reverse up to down, then swap the symmetry
 * 1 2 3     7 8 9     7 4 1
 * 4 5 6  => 4 5 6  => 8 5 2
 * 7 8 9     1 2 3     9 6 3
 */

func rotate_1(matrix [][]int)  {
	// 找规律
	reverse_nums(matrix)
	for i := 0; i < len(matrix); i++ {
		for j := i+1; j < len(matrix); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

func reverse_nums(matrix [][]int) {
	for i, j := 0, len(matrix)-1; i < j; i, j = i+1, j-1 {
		matrix[i], matrix[j] = matrix[j], matrix[i]
	}
}
