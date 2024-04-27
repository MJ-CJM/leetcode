// -*- coding:utf-8 -*-
// @Time : 2021/6/22 1:13 下午
// @Author: MJ-CJM
// @File : leetcode/6.22~ Offer 38. 字符串的排列
package main

import (
	"fmt"
	"strings"
)

func permutation(s string) []string {
	s_l := strings.Split(s, "")
	var res []string
	var tmp []string
	dfsPer(s_l, tmp, &res)
	return res
}

func dfsPer(l []string, tmp []string, res *[]string) {
	// terminator
	if len(l) == 0 {
		s1 := strings.Join(tmp, "")
		if !isInStr(s1, *res) {
			*res = append(*res, s1)
		}
		return
	}

	// process
	for i := 0; i < len(l); i++ {
		c1 := tmp
		c1 = append(c1, l[i])
		tmp2 := make([]string, len(l)-1)
		copy(tmp2[:i], l[:i])
		copy(tmp2[i:], l[i+1:])
		dfsPer(tmp2, c1, res)
	}
}

func isInStr(s string, l []string) bool {
	for _, v := range l {
		for i := 0; i < len(s); i++ {
			if s[i] != v[i] {
				break
			}
			if i == len(s)-1 {
				return true
			}
		}
	}
	return false
}

func main() {
	s := "abc"
	fmt.Println(permutation(s))
}
