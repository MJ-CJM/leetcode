package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 主函数，找到所有路径和等于 targetSum 的路径数目
func pathSum(root *TreeNode, targetSum int) int {
	prefixSumCount := map[int]int{0: 1}
	return dfs(root, 0, targetSum, prefixSumCount)
}

// 深度优先搜索
func dfs(node *TreeNode, currSum, targetSum int, prefixSumCount map[int]int) int {
	if node == nil {
		return 0
	}

	// 更新当前路径的前缀和
	currSum += node.Val

	// 获取到达当前节点的路径中，满足路径和等于 targetSum 的路径数
	res := prefixSumCount[currSum-targetSum]

	// 更新哈希表中的当前前缀和的次数
	prefixSumCount[currSum]++

	// 递归处理左右子树
	res += dfs(node.Left, currSum, targetSum, prefixSumCount)
	res += dfs(node.Right, currSum, targetSum, prefixSumCount)

	// 回溯，移除当前节点的前缀和数量
	prefixSumCount[currSum]--

	return res
}

