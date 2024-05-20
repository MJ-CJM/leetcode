// -*- coding:utf-8 -*-
// @Time : 2024/5/15 00:21
// @Author: MJ-CJM
// @File : leetcode/52.  n 皇后问题 2
package main

func totalNQueens(n int) int {
	var res int = 0
	var solue = make([]int, n)
	solveQueens(0, n, &solue, &res)
	return res
}

func solveQueens(row, n int, solve *[]int, res *int) {
	if row == n {
		*res++
		return
	}

	for col := 0; col < n; col++ {
		if isOk2(row, col, n, *solve) {
			(*solve)[row] = col
			solveQueens(row+1, n, solve, res)
		}
	}
}

func isOk2(row, col, n int, solve []int) bool {
	leftup, rightup := col - 1, col + 1
	for i := row - 1; i >= 0; i-- {
		if solve[i] == col {
			return false
		}
		if leftup >= 0 {
			if solve[i] == leftup {
				return false
			}
		}
		if rightup < n {
			if solve[i] == rightup {
				return false
			}
		}
		leftup--
		rightup++
	}
	return true
}
