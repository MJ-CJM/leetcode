// -*- coding:utf-8 -*-
// @Time : 2024/5/7 00:37
// @Author: MJ-CJM
// @File : leetcode/439.三元解析表达式
package main

func parseTernary(expression string) string {
	stack := []byte{}

	// 从右到左遍历表达式
	for i := len(expression) - 1; i >= 0; i-- {
		c := expression[i]
		if c == ':' {
			continue
		} else if c == '?' {
			i--
			if expression[i] == 'T' {
				tmp := stack[len(stack)-1]
				stack = stack[:len(stack)-2]
				stack = append(stack, tmp)
			} else {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, c)
		}

	}
	return string(stack[0])
}
