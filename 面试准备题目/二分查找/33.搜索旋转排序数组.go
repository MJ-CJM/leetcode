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
	if len(nums) == 0{
		return -1
	}
	left := 0
	right := len(nums) - 1
	for left <= right{
		mid := left + (right-left)/2
		if nums[mid] == target{
			return mid
		}
		if nums[0] <= nums[mid]{
			if nums[0] <= target && target < nums[mid]{
				right = mid - 1
			}else{
				left = mid + 1
			}
		}else{
			if nums[mid] < target && target <= nums[right]{
				left = mid + 1
			}else{
				right = mid - 1
			}
		}
	}
	return -1
}