// -*- coding:utf-8 -*-
// @Time : 2024/4/27 22:45
// @Author: MJ-CJM
// @File : leetcode/1056.易混淆数
package main

func confusingNumber(n int) bool {
	origin := n
	newN := 0

	for n > 0 {
		tmp := n % 10
		n /= 10

		switch tmp {
		case 0, 1, 8:
			newN = newN * 10 + tmp
		case 6:
			newN = newN * 10 + 9
		case 9:
			newN = newN * 10 + 6
		default:
			return false
		}
	}

	return newN != origin
}
