// -*- coding:utf-8 -*-
// @Time : 2024/5/5 23:17
// @Author: MJ-CJM
// @File : leetcode/252.会议室
package main


import "sort"
func canAttendMeetings(intervals [][]int) bool {
	n := len(intervals)
	if n == 0 {
		return true
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	for i := 0; i < n - 1; i++ {
		// 前者结束不能大于后者的开始
		if intervals[i][1] > intervals[i+1][0] {
			return false
		}
	}

	return true
}


