// -*- coding:utf-8 -*-
// @Time : 2024/5/10 23:13
// @Author: MJ-CJM
// @File : leetcode/5.最长回文子串
package main

// 中心扩展法
func longestPalindrome(s string) string {
	if s == "" {
		return ""
	}

	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		q := centerExpand(s, i, i)
		o := centerExpand(s, i, i + 1)

		maxl := 0
		if q > o {
			maxl = q
		} else {
			maxl = o
		}

		if maxl > end - start {
			// 正确得到起始位置
			start = i - (maxl - 1)/2
			end = i + maxl/2
		}
	}

	return s[start:end+1]
}

func centerExpand(s string, i, j int) int {
	for i > 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}

	return j - i - 1
}


/*
定义状态：

dp[i][j] 表示字符串 s 从索引 i 到索引 j 的子串是否是回文串。
状态转移方程：

dp[i][j] = (s[i] == s[j]) && (j - i < 3 || dp[i+1][j-1])
解释：
如果 s[i] == s[j] 且 j - i < 3（子串长度为2或3），那么 dp[i][j] = true。
如果 s[i] == s[j] 且 dp[i+1][j-1] == true，那么 dp[i][j] = true。
否则，dp[i][j] = false。
初始化：

所有长度为1的子串都是回文，即 dp[i][i] = true。
遍历顺序：

由于 dp[i][j] 的状态依赖于 dp[i+1][j-1]，我们需要从子串长度为2开始遍历，依次计算长度为3、4...直到 n 的子串。
记录最长回文子串：

在动态规划的过程中，记录最长的回文子串的起始索引和长度。
 */
// 动态规划
func longestPalindrome(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	// 初始化 dp 数组
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	start, maxLength := 0, 1

	// 每个单个字符都是回文
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}

	// 填充 dp 表
	for length := 2; length <= n; length++ {
		for i := 0; i < n-length+1; i++ {
			j := i + length - 1
			if s[i] == s[j] {
				if length == 2 {
					dp[i][j] = true
				} else {
					dp[i][j] = dp[i+1][j-1]
				}
				if dp[i][j] && length > maxLength {
					start = i
					maxLength = length
				}
			} else {
				dp[i][j] = false
			}
		}
	}

	return s[start : start+maxLength]
}

