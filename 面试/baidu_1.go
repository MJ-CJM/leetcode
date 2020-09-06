package main

import "fmt"

func main() {
	var n int
	var arr []int
	_, _ = fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var x int
		_, _ = fmt.Scan(&x)
		arr = append(arr, x)
	}
	res := solve(n, arr)
	fmt.Printf("%d", res)
}

func solve(n int, arr []int) int {
	if n <= 2 {
		return -1
	}
	res := 0
	for i := n; i >= 1; i-- {
		tmp := arr[:i]
		out := permuteUnique(tmp)
		if len(out) != 0 {
			for j := 0; j < len(out); j++ {
				linshi := count(out[j])
				if linshi % 90 == 0{
					if linshi > res {
						res = linshi
					}
				}
			}
		}
		if res != 0 {
			break
		}
	}
	return res
}

func count(nums []int) int {
	n := len(nums)
	res := 0
	for i := 0; i < n; i++ {
		res = res * 10 + nums[i]
	}
	return res
}


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
