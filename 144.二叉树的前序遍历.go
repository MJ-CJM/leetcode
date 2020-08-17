package main


// 迭代
func PreOrderTraversal(root *TreeNode) []int{
	if root == nil{
		return nil
	}
	if root.Right == nil && root.Left == nil{
		return []int{root.Val}
	}
	stack := []*TreeNode{root}
	res := []int{}
	for len(stack) != 0{
		tmp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, tmp.Val)
		if tmp.Right != nil{
			stack = append(stack, tmp.Right)
		}
		if tmp.Left != nil{
			stack = append(stack, tmp.Left)
		}
	}
	return res
}