// -*- coding:utf-8 -*-
// @Time : 2024/5/7 01:17
// @Author: MJ-CJM
// @File : leetcode/1420.第一个唯一数字
package main

import "sort"

// 暴力
type FirstUnique struct {
	q    []int
	mapQ map[int]int
}

func Constructor(nums []int) FirstUnique {
	n := len(nums)
	if n == 0 {
		return FirstUnique{
			q:    nums,
			mapQ: map[int]int{},
		}
	}

	tmpM := make(map[int]int)
	sort.Ints(nums)
	for i := 0; i < n; i++ {
		tmpM[nums[i]]++
	}

	return FirstUnique{
		q:    nums,
		mapQ: tmpM,
	}
}

func (this *FirstUnique) ShowFirstUnique() int {
	for i := 0; i < len(this.q); i++ {
		if this.mapQ[this.q[i]] == 1 {
			return this.q[i]
		}
	}
	return -1
}

func (this *FirstUnique) Add(value int) {
	this.q = append(this.q, value)
	this.mapQ[value]++
	sort.Ints(this.q)
}

// 优化后
type FirstUnique struct {
	q    []int
	mapQ map[int]int
	index   int
}

func Constructor(nums []int) FirstUnique {
	index := -1
	n := len(nums)
	if n == 0 {
		return FirstUnique{
			q:    nums,
			mapQ: map[int]int{},
			index: index,
		}
	}

	countM := make(map[int]int)
	for i := 0; i < n; i++ {
		countM[nums[i]]++
	}

	q := []int{}
	for i := 0; i < n; i++ {
		if countM[nums[i]] == 1 {
			q = append(q, nums[i])
			index = 0
		}
	}

	return FirstUnique{
		q:    q,
		mapQ: countM,
		index: index,
	}
}

func (this *FirstUnique) ShowFirstUnique() int {
	if this.index >= 0 && len(this.q) > 0 {
		for i := this.index; i < len(this.q); i++ {
			if this.mapQ[this.q[i]] == 1 {
				this.index = i
				return this.q[i]
			}
		}
		this.index = len(this.q) - 1
	}
	return -1
}

func (this *FirstUnique) Add(value int) {
	this.mapQ[value]++
	if this.mapQ[value] == 1 {
		this.q = append(this.q, value)
		if this.index < 0 {
			this.index = 0
		}
	}
	return
}
