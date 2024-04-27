// -*- coding:utf-8 -*-
// @Time : 2021/6/14 4:12 下午
// @Author: MJ-CJM
// @File : leetcode/6.14~374. 猜数字大小
package main

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is lower than the guess number
 *			      1 if num is higher than the guess number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	left := 0
	right := n
	for left <= right {
		mid := left + (right - left)>>1
		// tmp := guess(mid)
		tmp := mid
		if tmp == 1 {
			left = mid + 1
		} else if tmp == -1 {
			right = mid - 1
		} else {
			return mid
		}
	}
	return -1
}
