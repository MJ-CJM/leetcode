package main

import (
	"strconv"
)

// 全排列的解法
func getPermutation(n int, k int) string {
	var nums []string
	for i := 1; i <= n; i++ {
		nums = append(nums, strconv.Itoa(i))
	}
	n_nums := len(nums)
	level := 0
	result := []string{}
	iterm := ""
	dfs_string(nums, n_nums, level, iterm, &result)
	return result[k-1]
}

func dfs_string(nums []string, n int, level int, iterm string, result *[]string) {
	// terminator
	if level == n {
		tmp := iterm
		*result = append(*result, tmp)
		return
	}
	// process && drill down
	for i := 0; i < len(nums); i++ {
		c1 := iterm + nums[i]
		tmp1 := make([]string, len(nums)-1)
		copy(tmp1[:i], nums[:i])
		copy(tmp1[i:], nums[i+1:])
		dfs_string(tmp1, n, level+1, c1, result)
	}
}

// 数学解法
func getPermutation_1(n int, k int) string {
	factorial := make([]int, n)
	factorial[0] = 1
	for i := 1; i < n; i++ {
		factorial[i] = factorial[i-1] * i
	}
	k--

	ans := ""
	valid := make([]int, n+1)
	for i := 0; i < len(valid); i++ {
		valid[i] = 1
	}
	for i := 1; i <= n; i++ {
		order := k / factorial[n - i] + 1
		for j := 1; j <= n; j++ {
			order -= valid[j]
			if order == 0 {
				ans += strconv.Itoa(j)
				valid[j] = 0
				break
			}
		}
		k %= factorial[n-i]
	}
	return ans
}