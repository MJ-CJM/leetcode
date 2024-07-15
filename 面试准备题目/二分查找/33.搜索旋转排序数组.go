package main

//func search(nums []int, target int) int {
//	if len(nums) == 0{
//		return -1
//	}
//	left := 0
//	right := len(nums) - 1
//	for left < right{
//		mid := left + (right - left)/2
//		if nums[left] < nums[mid]{
//			if nums[left] > target || nums[mid] < target{
//				left = mid + 1
//			}else{
//				right = mid -1
//			}
//		}else{
//			if nums[mid] > target || nums[right] < target{
//				right = mid - 1
//			}else{
//				left = mid + 1
//			}
//		}
//	}
//	if nums[right] != target{
//		right = -1
//	}
//	return right
//}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		// 如果左半部分有序
		if nums[left] < nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[left] > nums[mid] { // 如果右半部分有序
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else { // 处理重复元素的情况
			left++
		}
	}

	return -1
}