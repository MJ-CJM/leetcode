package main

// 递归
func postorderTraversal(root *TreeNode) []int {
	if root == nil{
		return nil
	}
	var res []int
	if root.Left != nil{
		lres := postorderTraversal(root.Left)
		res = append(res, lres...)
	}
	if root.Right != nil{
		rres := postorderTraversal(root.Right)
		res = append(res, rres...)
	}
	res = append(res, root.Val)
	return res
}

// 迭代
func postorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	stack := []*TreeNode{root}
	for len(stack) > 0{
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if top.Left != nil{
			stack = append(stack, top.Left)
		}
		if top.Right != nil{
			stack = append(stack, top.Right)
		}
		res = append(res, top.Val)
	}
	return reverse(res)
}

func reverse(nums []int) []int {
	n := len(nums)
	if n <= 1{
		return nums
	}
	for i, j := 0, n-1; i < j; i,j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}