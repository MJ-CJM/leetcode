// -*- coding:utf-8 -*-
// @Time : 2024/5/17 22:22
// @Author: MJ-CJM
// @File : leetcode/122.买卖股票的最好时机2
package main

func maxProfit2(prices []int) int {
	res := 0
	n := len(prices)
	if n <= 1 {
		return 0
	}

	for i := 1; i < n; i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}
	return res
}