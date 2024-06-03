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
