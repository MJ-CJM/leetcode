package main

func exist(board [][]byte, word string) bool {
	n := len(board)
	m := len(board[0])
	e := len(word)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if search_word(board, i, j, 0, word, e) {
				return true
			}
		}
	}
	return false
}

func search_word(board [][]byte, i int, j int, level int, word string, e int) bool {
	// terminator
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
		return false
	}
	if board[i][j] != word[level] {
		return false
	}
	tmp := board[i][j]
	board[i][j] = ' '
	level++
	if level == e {
		return true
	}
	// process && drill down
	if search_word(board, i+1, j, level, word, e) {
		return true
	}
	if search_word(board, i-1, j, level, word, e) {
		return true
	}
	if search_word(board, i, j+1, level, word, e) {
		return true
	}
	if search_word(board, i, j-1, level, word, e) {
		return true
	}
	// reverse state
	board[i][j] = tmp
	return false
}

