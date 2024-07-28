package main

func jump(nums []int) int {
	n := len(nums)
	step := 0
	end := 0
	maxIndex := 0
	// i 跳到 n-2 位置，是为了防止在 n - 1 时，step 还要 ++
	for i := 0; i < n - 1; i++ {
		maxIndex = max(maxIndex, i + nums[i])
		if i == end {
			step++
			end = maxIndex
		}
	}
	return step
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
