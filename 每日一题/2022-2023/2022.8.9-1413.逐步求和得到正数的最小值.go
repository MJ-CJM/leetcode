// -*- coding:utf-8 -*-
// @Time : 2022/8/9 11:40 PM
// @Author: MJ-CJM
// @File : leetcode/2022.8.9-1413.逐步求和得到正数的最小值
package main

func minStartValue(nums []int) int {
	tmp := 0
	res := nums[0]
	for i, v := range nums {
		tmp += v
		nums[i] = tmp
		if tmp < res {
			res = tmp
		}
	}
	if res >= 0 {
		return 1
	} else {
		return -res + 1
	}

	//	 累加的思想
	//	presum, ans := 0, 1
	//	for _, num := range nums {
	//		presum += num
	//		if d := 1 - presum; d > ans {
	//			ans = d
	//		}
	//	}
	//	return ans
}
