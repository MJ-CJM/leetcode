package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	var result [][]int
	var iterm []int
	sort.Ints(candidates) // 排序
	dfsSum(candidates, target, 0, iterm, &result)
	return result
}

func dfsSum(candidates []int, target, start int, iterm []int, result *[][]int) {
	if target == 0 { // 如果当前和等于目标值，则找到一个解
		cur := make([]int, len(iterm))
		copy(cur, iterm)
		*result = append(*result, cur)
		return
	}

	for i := start; i < len(candidates); i++ {
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}
		if candidates[i] > target {
			break
		}
		iterm = append(iterm, candidates[i])
		dfsSum(candidates, target - candidates[i], i + 1, iterm, result)
		iterm = iterm[:len(iterm) - 1]
	}
}
