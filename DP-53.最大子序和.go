package main

// 暴力求解
//func maxsubarray(nums []int) int {
//	if nums == nil {
//		return 0
//	}
//	n := len(nums)
//	result := nums[0]
//	for i := 0; i < n; i++ {
//		for j := i; j < n; j++ {
//			index := i
//			temp := 0
//			for index <= j{
//				temp += nums[index]
//				index++
//			}
//			if temp > result{
//				result = temp
//			}
//		}
//	}
//	return result
//}

// 动态规划
// 重复子问题：max_sum(i) = Max(max_sum(i-1), 0) + a[i]
// 定义状态数组：f[i]
// 状态转移方程：f[i] = max(f[i-1],0) + a[i]
func maxSubArray(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	result := 0
	for i := 0; i < n; i++ {
		if i == 0{
			dp[i] = nums[i]
			result = dp[i]
		}else{
			dp[i] = max(dp[i-1]+nums[i], nums[i])
			if dp[i] > result{
				result = dp[i]
			}
		}
	}
	return result
}
