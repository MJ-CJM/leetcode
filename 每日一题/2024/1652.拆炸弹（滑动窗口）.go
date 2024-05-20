// -*- coding:utf-8 -*-
// @Time : 2024/5/5 20:17
// @Author: MJ-CJM
// @File : leetcode/1652.拆炸弹（滑动窗口）
package main

func decrypt(code []int, k int) []int {
	n := len(code)
	res := make([]int, n) // 初始化长度为 n 的结果切片

	if k == 0 {
		// 直接返回全 0 的数组
		return res
	}

	if k > 0 {
		// 计算第一个元素的和
		sum := 0
		for i := 1; i <= k; i++ {
			sum += code[i%n] // 确保不会超过数组界限
		}
		for j := 0; j < n; j++ {
			res[j] = sum
			sum -= code[(j+1)%n]      // 移除当前元素
			sum += code[(j+k+1)%n]    // 添加下一个 k 个元素
		}
	} else if k < 0 {
		// 转换 k 为正数，方便计算
		k = -k
		sum := 0
		for i := 0; i < k; i++ {
			sum += code[(n- i - 1)%n] // 计算前 k 个元素的和，逆序
		}
		for j := 0; j < n; j++ {
			res[j] = sum
			sum -= code[(n+j-k)%n]   // 移除当前元素
			sum += code[(n+j)%n]     // 添加下一个 k 个元素
		}
	}

	return res
}

