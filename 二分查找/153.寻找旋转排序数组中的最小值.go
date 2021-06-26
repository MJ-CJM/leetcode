package main

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)>>1
		//  假如 right和left都是一个很大的数，那么right + left会溢出，而(right - left)/2 + left 先做减法不会溢出。 所以好处是让pivot变量避免溢出
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = right - 1
		}
	}
	return nums[left]
}
