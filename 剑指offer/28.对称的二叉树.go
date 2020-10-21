package main

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dfs_duicheng(root.Left, root.Right)
}

func dfs_duicheng(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if (left == nil && right != nil) || (left != nil && right == nil) || (left.Val != right.Val) {
		return false
	}
	return dfs_duicheng(left.Left, right.Right) && dfs_duicheng(left.Right, right.Left)
}
