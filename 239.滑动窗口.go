package main

func maxSlidingWindow(nums []int, k int) []int {
	if nums == nil{
		return []int{}
	}
	res := []int{}
	win := []int{}
	for i, v := range nums{
		if i >= k && win[0] <= i-k{
			win = win[1:]
		}
		for len(win) > 0 && nums[win[len(win)-1]] < v{
			win = win[:len(win)-1]
		}
		win = append(win, i)
		if i >= k-1{
			res = append(res, nums[win[0]])
		}
	}
	return res
}
