package main

/**
 * 逆时针旋转矩阵
 * @param matrix int整型二维数组 输入矩阵
 * @return int整型二维数组
 */
func RotateMatrix( matrix [][]int ) [][]int {
	// write code here
	n := len(matrix)
	m := len(matrix[0])
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			tmp := j - m + 1
			if tmp < 0 {
				tmp = -tmp
			}
			res[tmp][i] = matrix[i][j]
		}
	}
	return res
}
