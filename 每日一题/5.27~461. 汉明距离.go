// -*- coding:utf-8 -*-
// @Time : 2021/5/27 10:18 上午
// @Author: MJ-CJM
// @File : leetcode/5.27~461. 汉明距离
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func hammingDistance(x int, y int) int {
	x_bin := convertToBin(x)
	y_bin := convertToBin(y)
	n := len(x_bin)
	m := len(y_bin)
	var l int = 0
	var b int = 0
	var res int = 0
	if n >= m {
		l = n
		b = n - m
		var tmp []string
		for i := 0; i < b; i++ {
			tmp = append(tmp, "0")
		}
		y_bin = append(tmp, y_bin...)
	} else {
		l = m
		b = m - n
		var tmp []string
		for i := 0; i < b; i++ {
			tmp = append(tmp, "0")
		}
		x_bin = append(tmp, x_bin...)
	}
	for i := 0; i < l; i++ {
		if x_bin[i] != y_bin[i] {
			res++
		}
	}
	return res
}

func convertToBin(num int) []string {
	str := ""

	if num == 0 {
		return []string{"0"}
	}
	for ; num > 0; num /= 2 {
		tmp := num % 2
		str = strconv.Itoa(tmp) + str
	}
	return strings.Split(str, "")
}

func main() {
	fmt.Println(convertToBin(5))
}
