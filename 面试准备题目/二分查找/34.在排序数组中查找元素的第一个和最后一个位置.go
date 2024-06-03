package main

func searchRange(nums []int, target int) []int {
	res := []int{}
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] == target {
			res = append(res, i)
			break
		}
		if nums[i] > target {
			break
		}
	}
	if len(res) != 1 {
		res = append(res, -1)
	}
	for j := n-1; j >= 0; j-- {
		if nums[j] == target {
			res = append(res, j)
			break
		}
		if nums[j] < target {
			break
		}
	}
	if len(res) == 1 {
		res = append(res, -1)
	}
	return res
}

func searchRange2(nums []int, target int) []int {
	n := len(nums)

	if n < 0 {
		return []int{-1, -1}
	}
	first := findFirstOccurrence(nums, target)
	end := findendOccurrence(nums, target)

	return []int{first, end}
}

func findFirstOccurrence(nums []int, target int) int {
	left, right := 0, len(nums) - 1
	result := -1

	for left <= right {
		mid := left + (right - left) / 2

		if nums[mid] == target {
			result = mid
			right = mid - 1 // 继续在左半部分查找
		} else if nums[mid] < target {
			left = mid + 1 // 在右半部分查找
		} else {
			right = mid - 1 // 在左半部分查找
		}
	}

	return result
}

func findendOccurrence(nums []int, target int) int {
	left, right := 0, len(nums) - 1
	result := -1

	for left <= right {
		mid := left + (right - left) / 2

		if nums[mid] == target {
			result = mid
			left = mid + 1 // 继续在左半部分查找
		} else if nums[mid] < target {
			left = mid + 1 // 在右半部分查找
		} else {
			right = mid - 1 // 在左半部分查找
		}
	}

	return result
}
