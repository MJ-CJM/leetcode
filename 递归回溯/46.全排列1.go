package main

import (
	"fmt"
)

func permute(nums []int) [][]int {
	n := len(nums)
	level := 0
	result := [][]int{}
	iterm := []int{}
	_dfs_all(nums, iterm, level, n, &result)
	return result
}

func _dfs_all(nums []int, iterm []int, level int, n int, result *[][]int) {
	// terminator
	if level == n{
		tmp := make([]int, n)
		copy(tmp, iterm)
		*result = append(*result, tmp)
		return
	}
	// process && drill down
	for i := 0; i < len(nums); i++ {
		c1 := append(iterm, nums[i])
		tmp := make([]int, len(nums)-1)
		copy(tmp[0:], nums[0:i])
		copy(tmp[i:], nums[i+1:])
		_dfs_all(tmp, c1, level+1, n, result)
	}
}

func main() {
	sum := []int{1, 2, 3}
	fmt.Println(permute(sum))
}