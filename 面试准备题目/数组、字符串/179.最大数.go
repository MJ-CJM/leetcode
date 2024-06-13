package main

import (
	"fmt"
	"sort"
	"strings"
)

func largestNumber(nums []int) string {
	// 转换数字为字符串
	numStrs := make([]string, len(nums))
	for i, num := range nums {
		numStrs[i] = fmt.Sprintf("%d", num)
	}

	sort.Slice(numStrs, func(i, j int) bool {
		return numStrs[i] + numStrs[j] > numStrs[j] + numStrs[i]
	})

	if numStrs[0] == "0" {
		return "0"
	}

	return strings.Join(numStrs, "")
}
