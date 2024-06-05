package main

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		tmp := []*TreeNode{}
		for _, v := range queue {
			if v.Left != nil {
				tmp = append(tmp, v.Left)
			}
			if v.Right != nil {
				tmp = append(tmp, v.Right)
			}
		}
		res++
		queue = tmp
	}
	return res
}