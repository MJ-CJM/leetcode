// -*- coding:utf-8 -*-
// @Time : 2021/5/26 3:08 下午
// @Author: MJ-CJM
// @File : leetcode/5.26~1190. 反转每对括号间的子串
package main

import (
	"fmt"
	"strings"
)

// 堆栈解法
func reverseParentheses(s string) string {
	str := strings.Split(s, "")
	n := len(str)
	var stack []string
	for i := 0; i < n; i++ {
		if str[i] != ")" {
			stack = append(stack, str[i])
		} else {
			var tmp []string
			for stack[len(stack)-1] != "(" {
				tmp = append(tmp, stack[len(stack)-1])
				// 切片传递
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
			//tmp = reverseStrList(tmp)
			stack = append(stack, tmp...)
		}
	}
	return strings.Join(stack, "")
}

func reverseStrList(str []string) []string {
	for i := 0; i < len(str)/2; i++ {
		str[i], str[len(str)-i-1] = str[len(str)-i-1], str[i]
	}
	return str
}

func reverseString(s string) string {
	var bytes []byte = []byte(s)
	for i := 0; i < len(s)/2; i++ {
		// 定义一个变量存放从后往前的值
		tmp := bytes[len(s)-i-1]
		// 从后往前的值跟从前往后的值调换
		bytes[len(s)-i-1] = bytes[i]
		// 从前往后的值跟从后往前的值进行调换
		bytes[i] = tmp
	}
	str := string(bytes)
	return str
}

func main() {
	//s := "abc"
	//fmt.Println(reflect.TypeOf(reverseString(s)))
	//str := []string{"s", "a", "e"}
	//fmt.Println(reverseStrList(str))
	//str = str[:len(str)-1]
	//fmt.Println(str)
	s_1 := "(abcd)"
	fmt.Println(reverseParentheses(s_1))
}
