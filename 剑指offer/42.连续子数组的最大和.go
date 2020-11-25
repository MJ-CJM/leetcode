package main

func maxSubArray(nums []int) int {
	if nums == nil {
		return 0
	}
	max_v, sum := nums[0], 0
	for i := 0; i < len(nums); i++ {
		if sum <= 0 {
			sum = nums[i]
		}else{
			sum += nums[i]
		}
		if sum > max_v {
			max_v = sum
		}
	}
	return max_v
}
