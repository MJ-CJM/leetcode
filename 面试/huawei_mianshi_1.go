package main

func Solve(nums [][]int) int {
	m := len(nums[0])
	count := 0
	for i := 0; i < m; i++ {
		count += nums[0][i]
	}
	count++
	n := len(nums)
	target := make([][]int, n)
	for i := 0; i < n; i++ {
		target[i] = make([]int, count)
	}
	for i := 0; i < n; i++ {
		target[i][0] = 1
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			tmp := process_nums(nums[i][j])
			target[i] = append(target[i], tmp...)
		}
	}
	res := 0
	for j := 0; j < count; j++ {
		tmp := 0
		for i := 0; i < n; i++ {
			if target[i][j] == 1 {
				tmp += 1
			}
		}
		if tmp > res {
			res = tmp
		}
	}
	return res
}

func process_nums(num int) []int {
	res := []int{}
	for i := 0; i < num; i++ {
		if i == num - 1 {
			res = append(res, 1)
		}else{
			res = append(res, 0)
		}
	}
	return res
}
