// -*- coding:utf-8 -*-
// @Time : 2021/6/21 1:36 上午
// @Author: MJ-CJM
// @File : leetcode/6.21~401. 二进制手表
package main

import (
	"fmt"
	"strconv"
)

// 回溯解法
func readBinaryWatch(turnedOn int) []string {
	var res []string
	level := 0
	var tmp [10]int
	dfsReadWatch(turnedOn, level, tmp, &res)
	return res
}

func dfsReadWatch(target, level int, tmp [10]int, res *[]string) {
	// terminator
	if level == 10 {
		if target == 0 {
			s1 := caculate(tmp)
			*res = append(*res, s1)
		}
		return
	}
	// process
	tmp[level] = 1
	c1 := tmp

	// drill down
	dfsReadWatch(target-1, level+1, c1, res)

	// reverse states
	tmp[level] = 0
	c2 := tmp
	dfsReadWatch(target, level+1, c2, res)

}

func caculate(tmp [10]int) string {
	var res string
	c1 := 0
	c2 := 0
	for i := 0; i < 10; i++ {
		if i == 0 && tmp[i] == 1 {
			c1 += 8
			continue
		}
		if i == 1 && tmp[i] == 1 {
			c1 += 4
			continue
		}
		if i == 2 && tmp[i] == 1 {
			c1 += 2
			continue
		}
		if i == 3 && tmp[i] == 1 {
			c1 += 1
			continue
		}
		if i == 4 && tmp[i] == 1 {
			c2 += 32
			continue
		}
		if i == 5 && tmp[i] == 1 {
			c2 += 16
			continue
		}
		if i == 6 && tmp[i] == 1 {
			c2 += 8
			continue
		}
		if i == 7 && tmp[i] == 1 {
			c2 += 4
			continue
		}
		if i == 8 && tmp[i] == 1 {
			c2 += 2
			continue
		}
		if i == 9 && tmp[i] == 1 {
			c2 += 1
			continue
		}
	}
	if c2 < 10 {
		if c2 == 0 {
			res = strconv.Itoa(c1) + ":" + "00"
		} else {
			res = strconv.Itoa(c1) + ":" + "0" + strconv.Itoa(c2)
		}
	} else {
		res = strconv.Itoa(c1) + ":" + strconv.Itoa(c2)
	}
	return res
}

func main() {
	tmp := [10]int{0, 1, 0, 0, 1, 1, 0}
	res := caculate(tmp)
	fmt.Println(res)
	fmt.Println(readBinaryWatch(1))
}