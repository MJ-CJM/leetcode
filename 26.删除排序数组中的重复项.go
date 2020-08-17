package main

func removeDuplicates(nums []int) int {
	count := 0
	for i := len(nums)-1; i > 0; i-- {
		if nums[i] == nums[i-1]{
			nums = append(nums[:i], nums[i+1:]...)
			count++
		}
	}
	return len(nums)
}
