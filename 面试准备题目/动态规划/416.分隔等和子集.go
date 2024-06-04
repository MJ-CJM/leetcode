package main

/*
边界条件：

如果数组的总和是奇数，那么不可能将其分割成两个子集，因为两个整数的和必须是偶数。
如果数组为空或只有一个元素（且不是 0），也无法分割。
问题转换：

计算数组的总和，如果总和是奇数，直接返回 false。
计算目标和 target，即为总和的一半。
转换为：能否在数组中找到一个子集，使得子集和为 target。
动态规划解决：

定义一个布尔数组 dp，其中 dp[i] 表示是否存在和为 i 的子集。
初始化 dp[0] 为 true，因为和为 0 的子集是存在的（空集）。
遍历数组中的每个数 num，从后向前更新 dp 数组：对于每个 i，如果 dp[i - num] 为 true，则 dp[i] 也为 true。
 */
func canPartition(nums []int) bool {
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	// 如果总和是奇数，不能分割成两个和相等的子集
	if totalSum%2 != 0 {
		return false
	}

	target := totalSum / 2
	dp := make([]bool, target+1)
	dp[0] = true

	for _, num := range nums {
		for i := target; i >= num; i-- {
			dp[i] = dp[i] || dp[i-num]
		}
	}

	return dp[target]
}
