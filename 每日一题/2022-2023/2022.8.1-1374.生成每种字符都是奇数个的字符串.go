// -*- coding:utf-8 -*-
// @Time : 2022/8/1 12:57 AM
// @Author: MJ-CJM
// @File : leetcode/2022.8.1-1374.生成每种字符都是奇数个的字符串
package main

import "strings"

func generateTheString(n int) string {
	if n == 1 {
		return "a"
	}
	if n%2 == 1 {
		return strings.Repeat("a", n)
	}
	return "a" + strings.Repeat("b", n-1)
}
