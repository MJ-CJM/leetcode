// -*- coding:utf-8 -*-
// @Time : 2024/5/7 00:38
// @Author: MJ-CJM
// @File : leetcode/484.寻找排列
package main

func findPermutation(s string) []int {
	n := len(s) + 1
	result := make([]int, 0, n)
	stack := make([]int, 0, n)

	// 我们需要遍历到s的长度+1，因为s比perm短1
	for i := 1; i <= len(s); i++ {
		stack = append(stack, i)
		if s[i-1] == 'I' {
			// 遇到'I'，输出栈中所有数字（形成递增序列）
			for len(stack) > 0 {
				result = append(result, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
		}
	}

	// 添加最后一个数字
	stack = append(stack, n)

	// 清空栈，可能是一系列'D'结束或单纯是结尾
	for len(stack) > 0 {
		result = append(result, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return result
}
