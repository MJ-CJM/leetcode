// -*- coding:utf-8 -*-
// @Time : 2023/11/23 00:32
// @Author: MJ-CJM
// @File : leetcode/1427.字符串的左右移
package main

func stringShift(s string, shift [][]int) string {
	sb := []byte(s)
	for _, sh := range shift {
		direction, amount := sh[0], sh[1]%len(sb) // 使用取模以防止移动超过字符串长度的情况
		if direction == 0 {                       // 左移
			sb = append(sb[amount:], sb[:amount]...)
		} else { // 右移
			sb = append(sb[len(sb)-amount:], sb[:len(sb)-amount]...)
		}
	}
	return string(sb)
}
