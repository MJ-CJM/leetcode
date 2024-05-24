// -*- coding:utf-8 -*-
// @Time : 2024/5/17 21:44
// @Author: MJ-CJM
// @File : leetcode/826.安排工作以达到最大收益
package main

import "sort"

type Node struct {
	dif   int
	value int
}

// 双指针
func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	n := len(difficulty)
	difPro := []Node{}
	for i := 0; i < n; i++ {
		difPro = append(difPro, Node{
			dif:   difficulty[i],
			value: profit[i],
		})
	}
	sort.Slice(difPro, func(i, j int) bool {
		return difPro[i].value < difPro[j].value
	})

	sort.Ints(worker)
	m := len(worker)
	res := 0
	i := n - 1
	j := m - 1
	for i >= 0 && j >= 0 {
		if difPro[i].dif <= worker[j] {
			res += difPro[i].value
			j--
		} else {
			i--
		}
	}
	return res
}