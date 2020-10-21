package main

func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
	iterm := []int{}
	dfs_lujinghe(root, sum, iterm, &res)
	return res
}

func dfs_lujinghe(root *TreeNode, sum int, iterm []int, res *[][]int) {
	// terminator
	if root == nil {
		return
	}
	iterm = append(iterm, root.Val)

	if root.Val == sum && root.Left == nil && root.Right == nil {
		tmp := make([]int, len(iterm))
		copy(tmp, iterm)
		*res = append(*res, tmp)
	}
	// process && drill down
	dfs_lujinghe(root.Left, sum-root.Val, iterm, res)
	dfs_lujinghe(root.Right, sum-root.Val, iterm, res)
	iterm = iterm[:len(iterm)-1]
}


