// -*- coding:utf-8 -*-
// @Time : 2021/6/19 12:09 上午
// @Author: MJ-CJM
// @File : leetcode/6.19~1239. 串联字符串的最大长度
package main

// 回溯求解最优
func maxLength(arr []string) int {
	n := len(arr)
	res := 0
	var iterm string
	i := 0
	dfs_maxLength(arr, iterm, n, i, &res)
	return res
}

func dfs_maxLength(arr []string, iterm string, n, i int, res *int) {
	// terminator
	if !check_iterm(iterm) {
		return
	}
	if i == n {
		tmp := len(iterm)
		if tmp > *res {
			*res = tmp
		}
		return
	}
	// process
	c1 := iterm
	c2 := iterm + arr[i]

	// drill down
	dfs_maxLength(arr, c1, n, i+1, res)
	dfs_maxLength(arr, c2, n, i+1, res)

	// 不用 reverse states
}

func check_iterm(iterm string) bool {
	itermMap := make(map[byte]bool)
	n := len(iterm)
	for i := 0; i < n; i++ {
		if itermMap[iterm[i]] == false {
			itermMap[iterm[i]] = true
		} else {
			return false
		}
	}
	return true
}
