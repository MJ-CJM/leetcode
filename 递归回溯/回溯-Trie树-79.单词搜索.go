package main

func exist(board [][]byte, word string) bool {
	n := len(board)
	m := len(board[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if backtrace_word(board, word, i, j, 0){
				return true
			}
		}
	}
	return false
}

func backtrace_word(board [][]byte, word string, i int, j int, k int) bool {
	// terminator
	if k == len(word) {
		return true
	}
	if i < 0 || j < 0 || i == len(board) || j == len(board[0]) {
		return false
	}
	if board[i][j] != word[k] {
		return false
	}

	// process and drill down
	tmp := board[i][j]
	board[i][j] = ' '
	if backtrace_word(board, word, i-1, j, k+1) {
		return true
	}
	if backtrace_word(board, word, i+1, j, k+1) {
		return true
	}
	if backtrace_word(board, word, i, j-1, k+1) {
		return true
	}
	if backtrace_word(board, word, i, j+1, k+1) {
		return true
	}
	board[i][j] = tmp
	return false
}


