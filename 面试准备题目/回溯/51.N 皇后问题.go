// -*- coding:utf-8 -*-
// @Time : 2024/5/14 22:59
// @Author: MJ-CJM
// @File : leetcode/51.N 皇后问题
package main

// 回溯法求解
func solveNQueens(n int) [][]string {
	var res = make([]int, n)
	var solutions = [][]string{}
	NQueens(0, n, &res, &solutions)
	return solutions
}


func NQueens(row int, n int, result *[]int, solutions *[][]string) {
	if row == n {
		*solutions = append(*solutions, PrintQ(n, *result))
		return
	}
	// 从列开始
	for i := 0; i < n; i++ {
		if isOk(row, i, n, result) {
			(*result)[row] = i
			NQueens(row+1, n, result, solutions)
		}
	}
}

func isOk(row, colum, n int, result *[]int) bool {
	leftup, rightup := colum - 1, colum + 1
	for i := row - 1; i >= 0; i-- {
		if (*result)[i] == colum {
			return false
		}
		if leftup >= 0 {
			if (*result)[i] == leftup {
				return false
			}
		}
		if rightup < n {
			if (*result)[i] == rightup {
				return false
			}
		}
		leftup--
		rightup++
	}
	return true
}

func PrintQ(n int, result []int) []string {
	res := []string{}
	for i := 0; i < n; i++ {
		row := ""
		for j := 0; j < n; j++ {
			if result[i] == j {
				row += "Q"
			} else {
				row += "."
			}
		}
		res = append(res, row)
	}
	return res
}