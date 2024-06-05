package main

import "math"

/*
递归计算每个节点的最大贡献值：

对于每个节点，我们需要计算以它为终点的最大路径和。
节点的最大贡献值等于节点值加上左右子树的最大贡献值中的较大者（如果子树贡献为负则不取）。
更新全局最大路径和：

对于每个节点，计算以该节点为最高点的路径和：即节点值加上左右子树的最大贡献值之和。
更新全局最大路径和。
递归函数的设计：

递归函数返回节点的最大贡献值。
在递归过程中更新全局最大路径和。

 */
func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	var maxGain func(node *TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// 递归计算左右子节点的最大贡献值
		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)

		// 当前节点的最大路径和
		priceNewPath := node.Val + leftGain + rightGain

		// 更新全局最大路径和
		maxSum = max(maxSum, priceNewPath)

		// 返回节点的最大贡献值
		return node.Val + max(leftGain, rightGain)
	}

	maxGain(root)
	return maxSum
}

