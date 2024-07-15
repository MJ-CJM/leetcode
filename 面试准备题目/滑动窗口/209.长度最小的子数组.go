package main

func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	i := 0
	j := 0
	sum := 0
	minLength := n + 1

	for j < n {
		sum += nums[j]
		for sum >= target {
			minLength = min(minLength, j-i+1)
			sum -= nums[i]
			i++
		}
		j++
	}

	if minLength == n + 1 {
		return 0
	}
	return minLength
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
