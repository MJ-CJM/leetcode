package main

import "strconv"

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
