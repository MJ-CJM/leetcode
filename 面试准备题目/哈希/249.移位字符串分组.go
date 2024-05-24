// -*- coding:utf-8 -*-
// @Time : 2024/5/5 21:03
// @Author: MJ-CJM
// @File : leetcode/249.移位字符串分组
package main

func groupStrings(strings []string) [][]string {
	res := make([][]string, 0)
	n := len(strings)

	for i := 0; i < n; i++ {
		tmp := strings[i]
		flag := false
		for j := 0; j < len(res); j++ {
			tB := []byte(tmp)
			jB := []byte(res[j][0])

			lt := len(tB)
			lj := len(jB)

			if lt != lj {
				continue
			}
			add := (tB[0] - jB[0] + 26) % 26
			count := 1
			for k := 1; k < lt; k++ {
				tmpA := (tB[k] - jB[k] + 26) % 26
				if tmpA == add {
					count++
				}
			}
			if count == lt {
				res[j] = append(res[j], tmp)
				flag = true
			}
		}
		if flag == false {
			res = append(res, []string{tmp})
		}
	}

	return res
}
