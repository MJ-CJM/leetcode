// -*- coding:utf-8 -*-
// @Time : 2021/6/15 11:32 下午
// @Author: MJ-CJM
// @File : leetcode/6.15~852. 山脉数组的峰顶索引
package main

func peakIndexInMountainArray(arr []int) int {
	n := len(arr)
	for i := 1; i < n-1; i++ {
		if arr[i-1] < arr[i] && arr[i] > arr[i+1] {
			return i
		}
	}
	return -1
}
