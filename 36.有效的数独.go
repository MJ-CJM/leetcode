package main

import "strconv"

func isValidSudoku(board [][]byte) bool {
	maprow := make([]map[int]int, 9)
	mapcol := make([]map[int]int, 9)
	mapbox := make([]map[int]int, 9)

	for i := 0; i < 9; i++ {
		maprow[i] = make(map[int]int)
		mapcol[i] = make(map[int]int)
		mapbox[i] = make(map[int]int)
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				tmp, _ := strconv.Atoi(string(board[i][j]))
				maprow[i][tmp]++
				mapcol[j][tmp]++
				mapbox[(i/3)*3+j/3][tmp]++
				if maprow[i][tmp] > 1 || mapcol[j][tmp] > 1 || mapbox[(i/3)*3+j/3][tmp] > 1 {
					return false
				}
			}
		}
	}
	return true
}
