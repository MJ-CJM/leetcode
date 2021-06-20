// -*- coding:utf-8 -*-
// @Time : 2021/6/19 12:09 上午
// @Author: MJ-CJM
// @File : leetcode/6.19~1239. 串联字符串的最大长度
package main

// 回溯求解最优
func maxLength(arr []string) int {
	n := len(arr)
	res, tmp := 0, 0
	var iterm string
	i := 0
	dfs_maxLength(arr, iterm, n, i, tmp, res)
	return res
}

func dfs_maxLength(arr []string, iterm string, n, i, tmp, res int) {
	// terminator
	if i == n {

	}
	// process

	// drill down

	// reverse states

}
