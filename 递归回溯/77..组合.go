package main

func combine(n int, k int) [][]int {
	res := [][]int{}
	level := 1
	iterm := []int{}
	digui_1(n , level, iterm, k, &res)
	return res
}

func digui_1(n int, level int, iterm []int, k int, res *[][]int) {
	// terminator
	if len(iterm) == k {
		tmp := make([]int, k)
		copy(tmp, iterm)
		*res = append(*res, tmp)
		return
	}
	// process && drill down
	for j := level; j <= n; j++ {
		c1 := append(iterm, j)
		digui_1(n, j+1, c1, k, res)
	}
}
