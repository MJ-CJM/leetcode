// -*- coding:utf-8 -*-
// @Time : 2024/5/10 23:13
// @Author: MJ-CJM
// @File : leetcode/5.最长回文子串
package main

// 中心扩展法
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}

	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		q := centerExpand(s, i, i)
		o := centerExpand(s, i, i + 1)

		maxl := 0
		if q > o {
			maxl = q
		} else {
			maxl = o
		}

		if maxl > end - start {
			// 正确得到起始位置
			start = i - (maxl - 1)/2
			end = i + maxl/2
		}
	}

	return s[start:end+1]
}

func centerExpand(s string, i, j int) int {
	for i > 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}

	return j - i - 1
}
