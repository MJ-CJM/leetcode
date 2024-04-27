// -*- coding:utf-8 -*-
// @Time : 2021/6/18 12:50 上午
// @Author: MJ-CJM
// @File : leetcode/6.18~483. 最小好进制
package main

import (
	"math"
	"strconv"
)

// 二进制
func smallestGoodBase(n string) string {
	num, _ := strconv.Atoi(n)
	for m := int(math.Log2(float64(num))); m >= 1; m-- {
		l, r := 2, int(math.Pow(float64(num), 1/float64(m)))+1
		for l < r {
			k := (l + r) / 2
			sum := 1
			for j := 0; j < m; j++ {
				sum = sum*k + 1
			}
			if sum == num {
				return strconv.Itoa(k)
			} else if sum < num {
				l = k + 1
			} else {
				r = k
			}
		}
	}
	return strconv.Itoa(num - 1)
}

// 官方解法

