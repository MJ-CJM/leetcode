// -*- coding:utf-8 -*-
// @Time : 2023/11/6 23:24
// @Author: MJ-CJM
// @File : leetcode/1768.交替合并字符串
package main

func mergeAlternately(word1 string, word2 string) string {
	len_1, len_2 := len(word1), len(word2)
	i, j := 0, 0
	res := make([]byte, len_1+len_2)

	for k := 0; k < len(res); {
		if i < len_1 {
			res[k] = word1[i]
			i++
			k++
		}
		if j < len_2 {
			res[k] = word2[j]
			j++
			k++
		}
	}

	return string(res)
}
