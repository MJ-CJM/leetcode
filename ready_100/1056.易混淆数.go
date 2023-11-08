// -*- coding:utf-8 -*-
// @Time : 2023/11/9 00:58
// @Author: MJ-CJM
// @File : leetcode/1056.易混淆数
package main

func confusingNumber(n int) bool {
	var original = n
	var res int

	for n > 0 {
		check := n % 10
		n /= 10

		switch check {
		case 0, 1, 8:
			res = res*10 + check
		case 6:
			res = res*10 + 9
		case 9:
			res = res*10 + 6
		default:
			// 2,3,4,5,7 返回 false
			return false
		}
	}

	return res != original
}
