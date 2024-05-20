// -*- coding:utf-8 -*-
// @Time : 2024/5/2 17:11
// @Author: MJ-CJM
// @File : leetcode/159.至少包含两个字符的最长子串
package main

func lengthOfLongestSubstringTwoDistinct(s string) int {
	res := 0
	sB := []byte(s)
	n := len(sB)
	sMap := make(map[byte]int)

	if n <= 2 {
		return n
	}

	for i, j := 0, 1; i < n && j < n; {
		if _, ok := sMap[sB[i]]; ok {
			sMap[sB[i]]++
		} else {
			sMap[sB[i]] = 1
		}
		if _, ok := sMap[sB[j]]; ok {
			sMap[sB[j]]++
		} else {
			sMap[sB[j]] = 1
		}
		if len(sMap) > 2 {
			count := j - i
			if count > res {
				res = count
			}
			sMap = map[byte]int{}
			i++
			j = i + 1
			continue
		}
		if j == n - 1 {
			count := n - i
			if count > res {
				res = count
			}
			sMap = map[byte]int{}
			i++
			j = i + 1
			continue
		}
		j++
	}

	return res
}

