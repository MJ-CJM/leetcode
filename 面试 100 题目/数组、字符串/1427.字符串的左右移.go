// -*- coding:utf-8 -*-
// @Time : 2024/4/27 23:07
// @Author: MJ-CJM
// @File : leetcode/1427.字符串的左右移
package main

func stringShift(s string, shift [][]int) string {
	n := len(shift)
	if n <= 0 {
		return s
	}

	res := []byte(s)
	l := len(res)
	for i := 0; i < n; i++ {
		j := shift[i][0]
		// 考虑移动的长度是否超过了队列长度
		k := shift[i][1] % l
		tmp := make([]byte, 0)

		if j == 0 {
			tmp = append(res[k:], res[:k]...)
		} else if j == 1 {
			tmp = append(res[l-k:], res[:l-k]...)
		}
		res = tmp
	}

	return string(res)
}

