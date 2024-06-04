package main

import "sort"

func merge2(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	n := len(intervals)
	if n == 1 {
		return intervals
	}

	res := [][]int{}
	cur := intervals[0]
	for i := 1; i < n; i++ {
		if intervals[i][0] <= cur[1] && intervals[i][1] > cur[1] {
			cur[1] = intervals[i][1]
			continue
		}

		if intervals[i][0] > cur[1] {
			res = append(res, cur)
			cur = intervals[i]
		}
	}
	res = append(res, cur)
	return res
}
