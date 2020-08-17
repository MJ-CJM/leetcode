package main

func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	copy(nums, append(nums[n-k:], nums[0:n-k]...))
}

