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

// 移位实现位计数
func hammingDistance_2(x int, y int) (ans int) {
	for s := x ^ y; s > 0; s >>= 1 {
		ans += s & 1
	}
	return
}

// 对于移位计数的优化，在方法二中，对于 s=(10001100)_2s=(10001100)
// 的情况，我们需要循环右移 88 次才能得到答案。而实际上如果我们可以跳过两个 11 之间的 00，直接对 11 进行计数，那么就只需要循环 33 次即可。
func hammingDistance_3(x int, y int) (ans int) {
	for s := x ^ y; s > 0; s &= s - 1 {
		ans++
	}
	return
}

func main() {
	fmt.Println(convertToBin(5))
}
