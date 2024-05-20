// -*- coding:utf-8 -*-
// @Time : 2024/4/29 00:28
// @Author: MJ-CJM
// @File : leetcode/4.28-1017.负二进制转换
package main

import (
	"fmt"
	"strings"
)

func baseNeg2(n int) string {
	if n == 0 {
		return "0"
	}
	var result []string
	for n != 0 {
		remainder := n % -2
		n = n / -2

		// 如果余数是负的，调整余数和商
		if remainder < 0 {
			remainder += 2
			n += 1
		}

		// 将余数添加到结果中（作为最低位）
		result = append([]string{fmt.Sprintf("%d", remainder)}, result...)
	}
	return strings.Join(result, "")
}


