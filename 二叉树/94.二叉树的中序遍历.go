package main

// 递归
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	res := []int{}
	res = inorderTraversal(root.Left)
	res = append(res, root.Val)
	res = append(res, inorderTraversal(root.Right)...)

	return res
}

// 迭代
func inorderTraversal2(root *TreeNode) []int{
	if root == nil{
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	stack := []*TreeNode{}
	p := root
	res := []int{}
	for p != nil || len(stack) != 0{
		for p != nil{
			stack = append(stack, p)
			p = p.Left
		}
		tmp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, tmp.Val)
		p = tmp.Right
	}
	return res
}


