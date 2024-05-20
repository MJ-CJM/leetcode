// -*- coding:utf-8 -*-
// @Time : 2024/5/5 23:27
// @Author: MJ-CJM
// @File : leetcode/253.会议室2
package main


import "sort"
func minMeetingRooms(intervals [][]int) int {
	nums := []int{}
	for _, v := range intervals {
		nums = append(nums, v[0] * 10 + 2)
		nums = append(nums, v[1] * 10 + 1)
	}
	sort.Ints(nums)

	tmp := 0
	res := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] % 10 == 2 {
			tmp++
		} else {
			tmp--
		}
		if tmp > res {
			res = tmp
		}
	}

	return res
}
