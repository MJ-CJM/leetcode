package main

import "fmt"

func nextPermutation(nums []int)  {
	n := len(nums)
	if n <= 1 {
		return
	}
	i, j, k := n-2, n-1, n-1
	for i > 0 && nums[i] >= nums[j] {
		i--
		j--
	}
	if i >= 0 {
		for nums[i] >= nums[k] {
			k--
		}
		swap_2(nums, i, k)
	}

	for i, j := j, n-1; i < j; i, j = i+1, j-1 {
		swap_2(nums, i, j)
	}
	return
}

func swap_2(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func main() {
	nums := []int{1, 2, 3}
	swap(nums, 0, 1)
	fmt.Println(nums)
}