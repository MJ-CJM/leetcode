// -*- coding:utf-8 -*-
// @Time : 2024/5/2 15:06
// @Author: MJ-CJM
// @File : leetcode/1055.形成字符串的最短路径
package main

func check(s []byte, f []byte) int {
	n := len(s)
	if n <= 0 {
		return 0
	}

	l := len(f)
	for i := 0; i < l; i++ {
		if f[i] == s[0] {
			tmp := 0
			for j:= i; j < l && tmp < n; j++ {
				if s[tmp] == f[j] {
					tmp++
				}
			}
			if tmp == n {
				return  n
			}
		}
	}
	return 0
}

func shortestWay(source string, target string) int {
	// 贪心求解
	sB := []byte(source)
	tB := []byte(target)

	nS := len(sB)
	nT := len(tB)

	res := 0
	if nS >= nT {
		count := 0
		start := 0
		flag := nT
		for start < flag {
			match := check(tB[start:flag], sB)
			if match > 0 {
				count += flag - start
				start = flag
				flag = nT
				res++
				continue
			}
			flag--
		}
		if count == nT {
			return res
		}
		return -1
	}

	count := 0
	start := 0
	flag := nS
	for start < flag {
		match := check(tB[start:flag], sB)
		if match > 0 {
			count += flag - start
			start = flag
			flag += nS
			if flag > nT {
				flag = nT
			}
			res++
			continue
		}
		flag--
	}

	if count == nT {
		return res
	}

	return -1;
}

