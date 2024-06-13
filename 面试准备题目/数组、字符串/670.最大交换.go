package main

import "strconv"

func maximumSwap(num int) int {
	// 将整数转换为字符数组
	numStr := []byte(strconv.Itoa(num))
	n := len(numStr)

	// 记录每个数字最后出现的位置
	last := make(map[byte]int)
	for i := 0; i < n; i++ {
		last[numStr[i]] = i
	}

	// 从左到右扫描，找出第一个可以交换的数字
	for i := 0; i < n; i++ {
		for d := byte('9'); d > numStr[i]; d-- {
			if last[d] > i {
				// 交换 numStr[i] 和 numStr[last[d]]
				numStr[i], numStr[last[d]] = numStr[last[d]], numStr[i]
				// 将字符数组转换回整数并返回
				result, _ := strconv.Atoi(string(numStr))
				return result
			}
		}
	}

	// 如果没有找到可以交换的数字，返回原数字
	return num
}

