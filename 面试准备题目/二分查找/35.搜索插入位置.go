package main

func searchInsert(nums []int, target int) int {
	n := len(nums)
	left := 0
	right := n - 1
	res := n
	for left <= right {
		mid := left + (right - left) >> 1
		if target <= nums[mid]  {
			res = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return res
}
