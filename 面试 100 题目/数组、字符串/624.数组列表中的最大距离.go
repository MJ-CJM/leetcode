// -*- coding:utf-8 -*-
// @Time : 2024/4/27 23:18
// @Author: MJ-CJM
// @File : leetcode/624.数组列表中的最大距离
package main

type node struct {
	val	 	int
	id  	int
}

func quickSort(list []node, low, high int) {
	if low > high {
		return
	}

	temp := list[low]
	i := low
	j := high
	for i < j {
		for list[j].val >= temp.val && i < j {
			j--
		}
		for list[i].val <= temp.val && i < j {
			i++
		}
		if i < j {
			list[i], list[j] = list[j], list[i]
		}
	}

	list[low] = list[i]
	list[i] = temp

	quickSort(list, low, i - 1)
	quickSort(list, i + 1, high)
	return
}

func QuickSort(list []node) []node {
	quickSort(list, 0, len(list)-1)
	return list // 这里返回list是为了符合函数签名，尽管排序已经是原地完成的
}

func absInt (a int, b int) int {
	if a > b {
		return a - b
	} else {
		return b - a
	}
}


func maxDistance(arrays [][]int) int {
	n := len(arrays)
	if n <= 0 {
		return 0
	}

	// 初始化最小值和最大值列表
	minList := make([]node, n)
	maxList := make([]node, n)

	for i := 0; i < n; i++ {
		minList[i] = node{val: arrays[i][0], id: i}
		maxList[i] = node{val: arrays[i][len(arrays[i])-1], id: i}
	}
	QuickSort(minList)
	QuickSort(maxList)

	// 使用标准库进行排序
	//sort.Slice(minList, func(i, j int) bool { return minList[i].val < minList[j].val })
	//sort.Slice(maxList, func(i, j int) bool { return maxList[i].val < maxList[j].val })


	if minList[0].id != maxList[len(maxList)-1].id {
		return absInt(minList[0].val, maxList[len(maxList)-1].val)
	} else {
		tmp1 := absInt(minList[0].val, maxList[len(maxList) - 2].val)
		tmp2 := absInt(minList[1].val, maxList[len(maxList)-1].val)
		if tmp1 > tmp2 {
			return tmp1
		} else {
			return tmp2
		}
	}
}


//func QuickSort(list []node) []node {
//	n := len(list)
//	if n <= 1 {
//		return list
//	} else {
//		mid := list[0]
//		left := []node{}
//		right := []node{}
//		for i := 1; i < n; i++ {
//			if list[i].val < mid.val {
//				left = append(left, list[i])
//			} else {
//				right = append(right, list[i])
//			}
//		}
//		return append(append(QuickSort(left), mid), QuickSort(right)...)
//	}
//}