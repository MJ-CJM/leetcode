package main

func levelOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	queue := []*TreeNode{root}
	res := []int{}
	for len(queue) != 0 {
		tmp := []int{}
		tmp2 := []*TreeNode{}
		for _, v := range queue {
			tmp = append(tmp, v.Val)
			if v.Left != nil {
				tmp2 = append(tmp2, v.Left)
			}
			if v.Right != nil {
				tmp2 = append(tmp2, v.Right)
			}
		}
		res = append(res, tmp...)
		queue = tmp2
	}
	return res
}
