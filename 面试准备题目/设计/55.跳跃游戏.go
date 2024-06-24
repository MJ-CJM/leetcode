package main

func canJump(nums []int) bool {
	if nums == nil{
		return false
	}
	n := len(nums)
	endnum := n-1
	for i := n-1; i >= 0; i --{
		if nums[i] + i >= endnum{
			endnum = i
		}
	}
	return endnum == 0
}

func canJump2(nums []int) bool {
	ability := nums[0]

	for i := 0; i < len(nums); i++ {
		if ability < nums[i] {
			ability = nums[i]
		}

		if ability <= 0 && i < len(nums) - 1 {
			return false
		}
		ability--
	}
	return true
}