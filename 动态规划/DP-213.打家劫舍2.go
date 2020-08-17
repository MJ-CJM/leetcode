package main

// 动态规划：重复子问题
// 定义状态：a[i]
// 状态转移方程：a[i] = max(a[i-1], a[i-2]+nums[i])
func rob3(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	n := len(nums)
	if n == 1{
		return nums[0]
	}
	dp1 := make([]int, n)
	dp2 := make([]int, n)
	dp1[0] = nums[0]
	dp1[1] = max_int(nums[0], nums[1])
	dp2[0] = 0
	dp2[1] = nums[1]
	for i := 2; i < n; i++ {
		dp1[i] = max_int(dp1[i-1], dp1[i-2]+nums[i])
		dp2[i] = max_int(dp2[i-1], dp2[i-2]+nums[i])
	}
	return max_int(dp1[n-2], dp2[n-1])
}
