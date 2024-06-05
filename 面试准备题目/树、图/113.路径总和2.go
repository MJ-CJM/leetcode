package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 主函数，找到所有路径
func pathSum(root *TreeNode, targetSum int) [][]int {
	var result [][]int
	var path []int
	dfs(root, targetSum, &path, &result)
	return result
}

// 深度优先搜索和回溯
func dfs(node *TreeNode, targetSum int, path *[]int, result *[][]int) {
	if node == nil {
		return
	}

	// 将当前节点加入路径
	*path = append(*path, node.Val)

	// 检查是否到达叶子节点并且路径和等于目标和
	if node.Left == nil && node.Right == nil && node.Val == targetSum {
		// 复制当前路径
		tempPath := make([]int, len(*path))
		copy(tempPath, *path)
		*result = append(*result, tempPath)
	}

	// 继续递归遍历左子树和右子树
	dfs(node.Left, targetSum-node.Val, path, result)
	dfs(node.Right, targetSum-node.Val, path, result)

	// 回溯，移除当前节点
	*path = (*path)[:len(*path)-1]
}
