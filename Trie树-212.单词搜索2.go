package main

type TrieNode struct {
	word string
	children [26]*TrieNode
}

// 1. 暴力 2. trie树
// trie树解法
// a. all words --> Trie构建起prefix
// b. board, DFS
func findWords(board [][]byte, words []string) []string {
	root := &TrieNode{}
	for  _, w := range words {
		node := root
		for _, c := range w {
			if node.children[c-'a'] == nil {
				node.children[c-'a'] = &TrieNode{}
			}
			node = node.children[c-'a']
		}
		node.word = w
	}
	result := make([]string, 0)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			dfs_word(i, j, board, root, &result)
		}
	}
	return result
}

func dfs_word(i int, j int, board [][]byte, node *TrieNode, result *[]string) {
	// terminator
	if i < 0 || j < 0 || i == len(board) || j == len(board[0]) {
		return
	}
	c := board[i][j]
	if c == '#' || node.children[c-'a'] == nil {
		return
	}
	node = node.children[c-'a']
	if node.word != ""{
		*result = append(*result, node.word)
		node.word = "" // 防止重复添加
	}
	board[i][j] = '#'
	dfs_word(i+1, j, board, node, result)
	dfs_word(i-1, j, board, node, result)
	dfs_word(i, j+1, board, node, result)
	dfs_word(i, j-1, board, node, result)
	board[i][j] = c
}


