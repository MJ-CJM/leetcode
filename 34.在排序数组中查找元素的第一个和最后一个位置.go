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
