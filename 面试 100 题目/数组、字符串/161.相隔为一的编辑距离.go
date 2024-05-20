// -*- coding:utf-8 -*-
// @Time : 2024/4/29 23:44
// @Author: MJ-CJM
// @File : leetcode/161.相隔为一的编辑距离
package main

func isOneEditDistance(s string, t string) bool {
	n, m := len(s), len(t)
	if abs(n-m) > 1 {
		return false
	}

	if n > m {
		return isOneEditDistance(t, s)
	}

	// 是否可以用删除和增加
	if n != m {
		for i := 0; i < n; i++ {
			if s[i] != t[i] {
				return s[i:] == t[i+1:]
			}
		}
		return true
	}

	// 是否可以用替换操作
	flag := false
	for i := 0; i <  n; i++ {
		if s[i] != t[i] {
			if flag {
				return false
			}
			flag = true
		}
	}

	return flag
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
