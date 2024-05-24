// -*- coding:utf-8 -*-
// @Time : 2024/5/14 22:53
// @Author: MJ-CJM
// @File : leetcode/2244.完成所有任务所需的最小轮数
package main

func minimumRounds(tasks []int) int {
	res := 0
	taskMap := make(map[int]int)

	for _, v := range tasks {
		taskMap[v]++
	}

	for _, v := range taskMap {
		tmp := countTask(v)
		if tmp == -1 {
			return -1
		}
		res += tmp
	}

	return res
}


func countTask(input int) int {
	res := 0
	if input == 1 {
		return -1
	}

	if input % 3 == 0 {
		return input / 3
	}

	if input == 2 {
		return 1
	}


	for input > 0 {
		input -= 2
		res += 1
		if input % 3 == 0 {
			return res + input / 3
		}
	}

	return res
}
