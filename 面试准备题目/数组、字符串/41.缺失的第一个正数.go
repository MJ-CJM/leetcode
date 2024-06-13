package main

/*
调整数组，使元素就位：

遍历数组，将每个数字 nums[i] 放到它应当在的位置，即 nums[nums[i] - 1]。
这个过程中，我们会持续交换元素，直到所有能放到正确位置的元素都被放到了正确的位置上。
查找第一个缺失的正整数：

再次遍历数组，如果 nums[i] 不是 i + 1，说明缺失的正整数就是 i + 1。
如果所有位置上的数字都正确，那么缺失的正整数就是 n + 1。
 */
func firstMissingPositive(nums []int) int {
	n := len(nums)

	for i := 0; i < n; {
		if nums[i] > 0 && nums[i] <= n && nums[i] != nums[nums[i] - 1]{
			nums[i], nums[nums[i] - 1] = nums[nums[i] - 1], nums[i]
		} else {
			i++
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] != i + 1 {
			return i + 1
		}
	}
	return n + 1
}
