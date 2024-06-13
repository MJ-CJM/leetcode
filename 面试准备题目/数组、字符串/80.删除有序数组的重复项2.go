package main

func removeDuplicates(nums []int) int {
	n := len(nums)
	if n <= 2 {
		return n
	}

	j := 2 // 从第三个元素开始
	for i := 2; i < n; i++ {
		// 如果当前元素与 nums[j-2] 不同，则可以保留
		if nums[i] != nums[j-2] {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}
