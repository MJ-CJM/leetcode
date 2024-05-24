// -*- coding:utf-8 -*-
// @Time : 2024/5/5 21:36
// @Author: MJ-CJM
// @File : leetcode/422.有效的单词方块
package 矩阵

func validWordSquare(words []string) bool {
	for i, word := range words {
		var j int
		for _, w := range words {
			if i < len(w) {
				if j >= len(word) {
					return false
				}

				// 内容是否相等
				if word[j] != w[i] {
					return false
				}
				j++
			}
		}

		// 长度是否相等
		if j != len(word) {
			return false
		}
	}


	return true
}

