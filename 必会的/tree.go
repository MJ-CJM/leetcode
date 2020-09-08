package main

// Definition for a binary tree node.
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 层次遍历
func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := make([][]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
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
		result = append(result, tmp)
		queue = tmp2
	}
	return result
}

// 前序遍历
func PreOrderTravel(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	if root.Left != nil && root.Right != nil {
		return []int{root.Val}
	}
	stack := []*TreeNode{root}
	result := []int{}
	for len(stack) != 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, cur.Val)
		if cur.Right != nil {
			stack = append(stack, cur.Right)
		}
		if cur.Left != nil {
			stack = append(stack, cur.Left)
		}
	}
	return result
}

// 中序遍历
// 递归
func InOrderTravel_1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	res := InOrderTravel_1(root.Left)
	res = append(res, root.Val)
	res = append(res, InOrderTravel_1(root.Right)...)
	return res
}

// 迭代
func InOrderTravel(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	stack := []*TreeNode{}
	p := root
	res := []int{}
	for p != nil || len(stack) != 0 {
		for p != nil {
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

// 后序遍历
// 递归
func PostOrderTravel_1(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	if root.Left != nil {
		lres := PostOrderTravel_1(root.Left)
		res = append(res, lres...)
	}
	if root.Right != nil {
		rres := PostOrderTravel_1(root.Right)
		res = append(res, rres...)
	}
	res = append(res, root.Val)
	return res
}

// 迭代
func PostOrderTravel(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := []*TreeNode{root}
	res := []int{}
	for len(stack) != 0 {
		tmp := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if tmp.Left != nil {
			stack = append(stack, tmp.Left)
		}
		if tmp.Right != nil {
			stack = append(stack, tmp.Right)
		}
		res = append(res, tmp.Val)
	}
	return reverse1(res)
}