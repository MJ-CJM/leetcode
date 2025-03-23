// -*- coding:utf-8 -*-
// @Time : 2024/5/20 23:44
// @Author: MJ-CJM
// @File : leetcode/76.最小覆盖子串
package main

func minWindow(s string, t string) string {
	// 定义计数器，用于记录 t 中每个字符的出现次数
	counter := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		counter[t[i]]++
	}

	// 定义窗口左右指针和最小子串的起始位置及长度
	left, right := 0, 0
	minLen := len(s) + 1
	start := 0

	// 定义计数器，用于记录窗口中包含的 t 中字符的出现次数
	count := 0

	// 开始滑动窗口算法
	for right < len(s) {
		// 如果当前字符在 t 中出现，则更新计数器
		if counter[s[right]] > 0 {
			count++
		}
		// 更新 counter
		counter[s[right]]--

		// 当窗口中包含了 t 中的所有字符时，移动左指针缩小窗口
		for count == len(t) {
			// 更新最小子串的起始位置和长度
			if right - left + 1 < minLen {
				minLen = right - left + 1
				 = left
			}
			// 如果左指针指向的字符在 t 中出现，则更新计数器
			if counter[s[left]] == 0 {
				count--
			}
			// 恢复 counter
			counter[s[left]]++
			// 移动左指针
			left++
		}
		// 移动右指针
		right++
	}

	// 如果最小子串的长度超过了 s 的长度，说明不存在符合条件的子串，返回空字符串
	if minLen > len(s) {
		return ""
	}
	// 返回最小子串
	return s[start : start + minLen]
}

