// -*- coding:utf-8 -*-
// @Time : 2024/5/4 00:58
// @Author: MJ-CJM
// @File : leetcode/找出变位映射
package main

func anagramMappings(nums1 []int, nums2 []int) []int {
	n := len(nums1)
	mapping := make([]int, 0)
	mapN := make(map[int]int)

	for i := 0; i < n; i++ {
		mapN[nums2[i]] = i
	}

	for i := 0; i < n; i++ {
		mapping = append(mapping, mapN[nums1[i]])
	}

	return mapping
}