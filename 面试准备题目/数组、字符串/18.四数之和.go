package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	n := len(nums)
	if n < 4 {
		return nil
	}
	sort.Ints(nums)
	res := [][]int{}
	for i := 0; i < n-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		for j := i + 1; j < n-2 && nums[i]+nums[j]+nums[i+1]+nums[i+1] <= target; j++ {
			if j > i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}
			for left, right := j+1, n-1; left < right; {
				if sum := nums[i] + nums[j] + nums[left] + nums[right]; sum == target {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					for left++; left < right && nums[left] == nums[left-1]; left++ {
					}
					for right--; right > left && nums[right] == nums[right+1]; right--{
					}
				}else if sum < target {
					left++
				}else{
					right--
				}
			}
		}
	}
	return res
}
