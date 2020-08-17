package main

func solveSudoku(board [][]byte)  {
	row := [9][9]bool{}
	col := [9][9]bool{}
	box := [9][9]bool{}
	for rowk := 0; rowk < 9; rowk++ {
		for colk := 0; colk < 9; colk++ {
			if board[rowk][colk] != '.' {
				num := board[rowk][colk] - '1'
				boxk := (rowk/3)*3 + colk/3
				row[rowk][num], col[colk][num], box[boxk][num] = true, true, true
			}
		}
	}
	fill(board, 0, row, col, box)
}

func fill(board [][]byte, n int, row [9][9]bool, col [9][9]bool, box [9][9]bool) bool {
	// terminator
	if n == 81 {
		return true
	}
	rowk := n / 9
	colk := n % 9
	if board[rowk][colk] != '.' {
		return fill(board, n+1, row, col, box)
	}
	boxk := (rowk/3)*3 + colk/3
	for num := 0; num < 9; num++ {
		if !row[rowk][num] && !col[colk][num] && !box[boxk][num] {
			board[rowk][colk] = byte('1'+num)
			row[rowk][num], col[colk][num], box[boxk][num] = true, true, true
			if fill(board, n+1, row, col, box) {
				return true
			}
			row[rowk][num], col[colk][num], box[boxk][num] = false, false, false
		}
	}
	board[rowk][colk] = '.'
	return false
}
