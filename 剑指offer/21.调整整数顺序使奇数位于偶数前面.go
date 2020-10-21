package main

func exchange(nums []int) []int {
	n := len(nums)
	i := 0
	j := n-1
	for i < j {
		if nums[i] % 2 == 1 && nums[j] % 2 == 0{
			i++
			j--
			continue
		}
		if nums[i] % 2 == 0 && nums[j] % 2 == 1 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
			continue
		}
		if nums[i] % 2 == 0 && nums[j] % 2 == 0 {
			j--
			continue
		}
		if nums[i] % 2 == 1 && nums[j] % 2 == 1 {
			i++
			continue
		}
	}
	return nums
}
