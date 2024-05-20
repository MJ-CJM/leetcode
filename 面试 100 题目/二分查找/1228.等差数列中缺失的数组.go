// -*- coding:utf-8 -*-
// @Time : 2024/5/8 23:05
// @Author: MJ-CJM
// @File : leetcode/1228.等差数列中缺失的数组
package main

func missingNumber(arr []int) int {
	n := len(arr)

	diff := (arr[n-1]-arr[0]) / n

	left := 0
	right := len(arr) -1

	for left < right - 1 {
		mid := left + (right - left) / 2
		if arr[mid] == arr[left] + diff * (mid - left) {
			left = mid
		} else {
			right = mid
		}
	}
	return arr[left] + diff
}
