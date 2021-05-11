package main

import (
	"strconv"
)

// 数学解法
func findNthDigit(n int) int {
	var digst = 1
	var count = 9
	var start = 1
	for n - count > 0 {
		n -= count
		digst += 1
		start *= 10
		count = digst * start * 9
	}
	var num = start + (n-1)/digst
	var res = strconv.Itoa(num)[(n-1)%digst] - '0'
	return int(res)
}