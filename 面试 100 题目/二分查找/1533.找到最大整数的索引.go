// -*- coding:utf-8 -*-
// @Time : 2024/5/8 23:45
// @Author: MJ-CJM
// @File : leetcode/1533.找到最大整数的索引
package main

func getIndex(reader *ArrayReader) int {
	left := 0
	right := reader.length() - 1

	for left < right {
		mid := left + (right - left) / 2

		if (right - left + 1) % 2 == 1 {
			// 总共奇数个
			flag := reader.compareSub(left, mid - 1, mid + 1, right)
			if flag == 0 {
				return mid
			} else if flag == 1 {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// 总共偶数个
			flag := reader.compareSub(left, mid, mid + 1, right)
			if flag == 0 {
				return mid
			} else if flag == 1 {
				right = mid
			} else {
				left = mid + 1
			}
		}
	}

	return left
}
