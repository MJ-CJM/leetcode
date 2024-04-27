// -*- coding:utf-8 -*-
// @Time : 2021/5/24 4:07 下午
// @Author: MJ-CJM
// @File : leetcode/5.24~664. 奇怪的打印机
package main

import "fmt"

func minChanges(a []int, k int) (ans int) {
	cnt := make([][1024]int, k)
	n := len(a)
	minCnt := n
	for i := 0; i < k; i++ {
		maxCnt := 0
		for j := i; j < n; j += k {
			cnt[i][a[j]]++
			maxCnt = max(maxCnt, cnt[i][a[j]])
		}
		ans += maxCnt
		minCnt = min(minCnt, maxCnt)
	}
	ans -= minCnt

	dp := make([][1024]int, k)
	for i := range dp {
		for j := 0; j < 1024; j++ {
			dp[i][j] = -1
		}
	}
	var f func(int, int) int
	f = func(p, xor int) (res int) {
		if p == 0 {
			return cnt[0][xor]
		}
		dv := &dp[p][xor]
		if *dv >= 0 {
			return *dv
		}
		defer func() { *dv = res }()
		for i := p; i < n; i += k {
			res = max(res, cnt[p][a[i]]+f(p-1, xor^a[i]))
		}
		return
	}
	return n - max(ans, f(k-1, 0))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println()
}
