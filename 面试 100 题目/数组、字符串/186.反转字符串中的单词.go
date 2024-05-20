// -*- coding:utf-8 -*-
// @Time : 2024/5/1 01:06
// @Author: MJ-CJM
// @File : leetcode/186.反转字符串中的单词
package main

func reverseWords(s []byte)  {
	// 双指针移动，先反转整个字符串，再逐个翻转
	n := len(s)

	for i, j := 0, n - 1; i < j; i, j = i + 1, j - 1{
		s[i], s[j] = s[j], s[i]
	}

	start := 0
	for i := 0; i <= n; i++ {
		if i == n || s[i] == ' '  {
			for j, k := start, i - 1; j < k; j, k = j + 1, k - 1 {
				s[j], s[k] = s[k], s[j]
			}
			start = i + 1
		}
	}
	return
}