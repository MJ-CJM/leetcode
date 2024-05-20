// -*- coding:utf-8 -*-
// @Time : 2024/5/20 00:28
// @Author: MJ-CJM
// @File : leetcode/547.省份数量
package main

func findCircleNum(M [][]int) int {
	n := len(M)
	if n == 0 {
		return n
	}
	union := NewUnionSet(n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if M[i][j] == 1 {
				union.Union(i, j)
			}
		}
	}
	return union.count
}

type UnionSet struct {
	count  int
	parent []int
}

func NewUnionSet(n int) *UnionSet {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UnionSet{
		count: n,
		parent: parent,
	}
}

func (u *UnionSet) Union(i, j int) {
	ri := u.find(i)
	rj := u.find(j)
	if ri != rj {
		u.parent[ri] = rj
		u.count--
	}
}

func (u *UnionSet) find(i int) int {
	root := i
	for u.parent[root] != root {
		root = u.parent[i]
	}
	for u.parent[i] != i { // 路径压缩
		i, u.parent[i] = u.parent[i], root
	}
	return root
}