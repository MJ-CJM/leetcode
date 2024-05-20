// -*- coding:utf-8 -*-
// @Time : 2024/4/29 23:33
// @Author: MJ-CJM
// @File : leetcode/280.摆动排序
package main

func wiggleSort(nums []int)  {
	n := len(nums);
	if n <= 1 {
		return
	}

	for i := 1; i < n; i++ {
		if i % 2 == 0 {
			if nums[i] > nums[i-1] {
				nums[i], nums[i-1] = nums[i-1], nums[i]
			}
		} else if i % 2 == 1 {
			if nums[i] < nums[i-1] {
				nums[i], nums[i-1] = nums[i-1], nums[i]
			}
		}
	}

	return
}
