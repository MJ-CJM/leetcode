// -*- coding:utf-8 -*-
// @Time : 2024/5/9 01:09
// @Author: MJ-CJM
// @File : leetcode/1231.分享巧克力
package main

func countSweets(nums []int, target int, k int) bool {
	count := 0
	tmp := 0
	for _, v := range nums {
		tmp += v
		if tmp > target {
			count++
			tmp = 0
		}
	}
	if count >= k + 1 {
		return true
	} else {
		return false
	}
}


func maximizeSweetness(sweetness []int, k int) int {
	low, high := 1, 0
	sum_sweet := 0
	for _, v := range sweetness {
		sum_sweet += v
	}
	high = sum_sweet / (k+1)

	for low <= high {
		mid := low + (high - low) / 2
		flag := countSweets(sweetness, mid, k)
		if flag {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}
