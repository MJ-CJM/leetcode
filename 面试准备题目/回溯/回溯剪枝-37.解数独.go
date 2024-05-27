package main

func solveSudoku(board [][]byte) {
	row := [9][9]bool{}  // 记录每行的数字使用情况
	col := [9][9]bool{}  // 记录每列的数字使用情况
	box := [9][9]bool{}  // 记录每个 3x3 宫格的数字使用情况

	// 初始化辅助结构
	for rowk := 0; rowk < 9; rowk++ {
		for colk := 0; colk < 9; colk++ {
			if board[rowk][colk] != '.' {
				num := board[rowk][colk] - '1'
				boxk := (rowk/3)*3 + colk/3
				row[rowk][num], col[colk][num], box[boxk][num] = true, true, true
			}
		}
	}
	// 开始填充数独
	fill(board, 0, row, col, box)
}

func fill(board [][]byte, n int, row [9][9]bool, col [9][9]bool, box [9][9]bool) bool {
	// terminator: 当所有格子都填完时，返回 true 表示成功
	if n == 81 {
		return true
	}
	rowk := n / 9   // 当前行
	colk := n % 9   // 当前列

	// 如果当前格子已经有数字，跳到下一个格子
	if board[rowk][colk] != '.' {
		return fill(board, n+1, row, col, box)
	}

	boxk := (rowk/3)*3 + colk/3  // 当前格子所在的 3x3 宫格编号
	for num := 0; num < 9; num++ {
		// 检查当前数字是否可以放在当前格子
		if !row[rowk][num] && !col[colk][num] && !box[boxk][num] {
			// 放置数字
			board[rowk][colk] = byte('1' + num)
			row[rowk][num], col[colk][num], box[boxk][num] = true, true, true

			// 递归处理下一个格子
			if fill(board, n+1, row, col, box) {
				return true
			}

			// 回溯：撤销放置
			row[rowk][num], col[colk][num], box[boxk][num] = false, false, false
			board[rowk][colk] = '.'
		}
	}
	// 如果所有数字都不能放置，返回 false
	return false
}

