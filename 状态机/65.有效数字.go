// -*- coding:utf-8 -*-
// @Time : 2021/6/19 4:22 下午
// @Author: MJ-CJM
// @File : leetcode/65.有效数字
package main

import "strings"

var (
	blank  = 0 // 空格
	digit1 = 1 // 数字（0-9）无前缀
	sign1  = 2 // +/- 无 e 前缀
	point  = 4 // '.'
	digit2 = 5 // 数字（0-9）有符号前缀
	e      = 6 // 'e'
	sign2  = 7 // +/- 有 e 前缀
	digit3 = 8 // 数字（0-9） 有 前缀
)

func isNumber_2(s string) bool {
	s = strings.TrimRight(s, " ")
	dfa := [][]int{
		[]int{blank, digit1, sign1, point, -1},
		[]int{-1, digit1, -1, digit2, e},
		[]int{-1, digit1, -1, point, -1},
		[]int{-1, digit2, -1, -1, e},
		[]int{-1, digit2, -1, -1, -1},
		[]int{-1, digit2, -1, -1, e},
		[]int{-1, digit3, -1, -1, -1},
		[]int{-1, digit3, -1, -1, -1},
	}

	state := 0 // blank start
	for i := 0; i < len(s); i++ {
		var newState int
		switch s[i] {
		case ' ':
			newState = 0
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			newState = 1
		case '+', '-':
			newState = 2
		case '.':
			newState = 3
		case 'e':
			newState = 4
		default:
			return false
		}
		state = dfa[state][newState]
		if state == -1 {
			return false
		}
	}
	return state == digit1 || state == digit2 || state == digit3
}
