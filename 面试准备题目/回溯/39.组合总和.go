package main

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	res := [][]int{}
	iterm := []int{}
	zuhe(candidates, iterm, target, 0, &res)
	return res
}

func zuhe(candidates []int, iterm []int, target int, left int, res *[][]int) {
	// terminator
	if target == 0 {
		tmp := make([]int, len(iterm))
		copy(tmp, iterm)
		*res = append(*res, tmp)
		return
	}
	// process && drill down
	for i := left; i < len(candidates); i++ {
		if target < candidates[i] {
			return
		}
		zuhe(candidates, append(iterm, candidates[i]), target-candidates[i], i, res)
	}
}

