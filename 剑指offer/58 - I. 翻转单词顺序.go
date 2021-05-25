// -*- coding:utf-8 -*-
// @Time : 2021/5/17 12:29 下午
// @Author: MJ-CJM
// @File : leetcode/58 - I. 翻转单词顺序
package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	str := strings.Split(s, " ")
	var res_s []string
	for i := len(str)-1; i >= 0; i-- {
		tmp := strings.TrimSpace(str[i])
		if len(tmp) > 0 {
			res_s = append(res_s, str[i])
		}
	}
	return strings.Join(res_s, " ")
}

// 无法输入连续字符串
func main() {
	var c byte
	var err error
	var d []string
	for ; err == nil; {
		_, err = fmt.Scanf("%c", &c)

		str := string(c)
		if str != "\n" {
			d = append(d, str)
		} else {
			break;
		}
	}
	s := strings.Join(d, " ")
	fmt.Println(s)
	res := reverseWords(s)
	fmt.Println(res)
}