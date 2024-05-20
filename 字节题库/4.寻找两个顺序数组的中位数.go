// -*- coding:utf-8 -*-
// @Time : 2024/5/10 21:12
// @Author: MJ-CJM
// @File : leetcode/寻找两个顺序数组的中位数
package main

// 暴力解法
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	res := []int{}
	i, j := 0, 0
	target := 0.0

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] <= nums2[j] {
			res = append(res, nums1[i])
			i++
		} else {
			res = append(res, nums2[j])
			j++
		}
	}

	if i < len(nums1) {
		res = append(res, nums1[i:]...)
	}

	if j < len(nums2) {
		res = append(res, nums2[j:]...)
	}

	n := len(res)
	if n % 2 == 1 {
		target = float64(res[n/2])
	} else {
		target = float64((res[n/2] + res[n/2-1]))/2
	}
	return target
}

// 二分查找第 k 小数求解
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		minIndex := totalLength / 2
		return float64(findKMinValue(nums1, nums2, minIndex + 1))
	} else {
		minIndex1 := totalLength / 2
		minIndex2 := totalLength / 2 - 1
		return float64(findKMinValue(nums1, nums2, minIndex1 + 1) + findKMinValue(nums1, nums2, minIndex2 + 1)) / 2.0
	}
	return 0
}

func findKMinValue(nums1, nums2 []int, k int) int {
	index1 := 0
	index2 := 0
	for {
		if index1 == len(nums1) {
			return nums2[index2 + k - 1]
		}

		if index2 == len(nums2) {
			return nums1[index1 + k - 1]
		}

		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}
		half := k / 2
		newIndex1 := min(index1 + half, len(nums1)) - 1
		newIndex2 := min(index2 + half, len(nums2)) - 1
		if nums1[newIndex1] <= nums2[newIndex2] {
			k -= (newIndex1 - index1 + 1)
			index1 = newIndex1 + 1
		} else {
			k -= (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}

	}

	return 0
}


func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}



