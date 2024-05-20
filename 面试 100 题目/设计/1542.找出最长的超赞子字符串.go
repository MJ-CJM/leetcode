// -*- coding:utf-8 -*-
// @Time : 2024/5/20 17:04
// @Author: MJ-CJM
// @File : leetcode/1542.找出最长的超赞子字符串
package main

/*
这段代码的基本思路是利用位运算来跟踪字符出现次数的奇偶性，并使用哈希表来记录不同状态的最早出现位置，从而快速判断某个子字符串是否可以重排成回文字符串。回文字符串要求最多只有一个字符出现奇数次，其余字符都出现偶数次。
 */
func longestAwesome(s string) int {
	// 定义哈希表，键为 mask，值为该 mask 出现的最早位置
	maskMap := make(map[int]int)
	// 初始化 maskMap，将 mask 0（空字符串）的位置设为 -1
	maskMap[0] = -1

	// 初始化结果为 0，当前 mask 为 0
	res, mask := 0, 0

	// 遍历字符串 s
	for i, char := range s {
		// 计算当前字符的奇偶性
		digit := int(char - '0')
		// 切换当前字符的奇偶性
		mask ^= 1 << digit

		// 如果 mask 已经出现过，说明找到一个超赞子字符串
		if idx, exists := maskMap[mask]; exists {
			// 更新结果为当前位置与之前 mask 出现的位置的差值和之前的最大结果中的最大值
			res = max(res, i - idx)
		} else {
			// 如果 mask 没有出现过，将当前 mask 的位置记录下来
			maskMap[mask] = i
		}

		// 遍历所有奇偶性情况下的 mask，检查是否存在能够交换字符构成回文字符串的情况
		for j := 0; j < 10; j++ {
			altMask := mask ^ (1 << j)
			if idx, exists := maskMap[altMask]; exists {
				// 更新结果为当前位置与之前 mask 出现的位置的差值和之前的最大结果中的最大值
				res = max(res, i - idx)
			}
		}
	}

	return res
}