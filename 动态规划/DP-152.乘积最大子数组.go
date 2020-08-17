package main

// 动态规划
// 重复子问题: max[i] = max(max[i-1]*a[i], a[i])
// 定义状态数组: f[i]
// 状态转移方程:  f[i] = max(f[i-1]*a[i], a[i])
func maxProduct(nums []int) int {
	n := len(nums)
	dp1 := make([]int, n)
	dp2 := make([]int, n)
	result := 0
	for i := 0; i < n; i++ {
		if i == 0 {
			dp1[i] = nums[i]
			dp2[i] = nums[i]
			result = dp1[i]
		} else {
			dp1[i] = min_3(nums[i], dp1[i-1]*nums[i], dp2[i-1]*nums[i])
			dp2[i] = max_3(nums[i], dp1[i-1]*nums[i], dp2[i-1]*nums[i])
			result = max_3(dp1[i], dp2[i], result)
		}
	}
	return result
}

func min_3(x, y, z int) int {
	result := x
	if y < result {
		result = y
	}
	if z < result {
		result = z
	}
	return result
}

func max_3(x, y, z int) int {
	result := x
	if y > result {
		result = y
	}
	if z > result {
		result = z
	}
	return result
}
