package main

import "fmt"

func permuteUnique(nums []int) [][]int {
	n := len(nums)
	level := 0
	result := [][]int{}
	iterm := []int{}
	_dfs_all2(nums, iterm, level, n, &result)
	return result
}

func _dfs_all2(nums []int, iterm []int, level int, n int, result *[][]int) {
	// terminator
	if level == n{
		if !_is_in(*result, iterm){
			tmp1 := make([]int, len(iterm))
			copy(tmp1, iterm)
			*result = append(*result, tmp1)
		}
		return
	}
	// process && drill dowm
	for i := 0; i < len(nums); i++ {
		c1 := append(iterm, nums[i])
		tmp := make([]int, len(nums)-1)
		copy(tmp[0:], nums[:i])
		copy(tmp[i:], nums[i+1:])
		_dfs_all2(tmp, c1, level+1, n, result)
	}
}

func _is_in(result [][]int, iterm []int) bool{
	for _, v := range result{
		for i := 0; i < len(v); i++ {
			if v[i] != iterm[i]{
				break
			}
			if i == len(v) - 1{
				return true
			}
		}
	}
	return false
}

func main() {
	sum := []int{1, 2, 3}
	fmt.Println(permuteUnique(sum))
}