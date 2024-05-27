package main

/*
给你一个整数数组 nums，返回 数组 answer ，其中 answer[i] 等于 nums 中除 nums[i] 之外其余各元素的乘积 。

题目数据 保证 数组 nums之中任意元素的全部前缀元素和后缀的乘积都在  32 位 整数范围内。

请 不要使用除法，且在 O(n) 时间复杂度内完成此题。
 */
// 先算左边乘积，再算右边乘积
func productExceptSelf(nums []int) []int {
	n := len(nums)
	left := make([]int, n)
	left[0] = 1
	for i := 1; i < n; i++ {
		left[i] = left[i-1] * nums[i-1]
	}

	right := 1
	for i := n - 1; i >= 0; i-- {
		left[i] = right * left[i]
		right *= nums[i]
	}
	return left
}