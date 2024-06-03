package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root == nil{
		return true
	}

	p, q := root, root
	stack1 := []*TreeNode{}
	stack2 := []*TreeNode{}

	for len(stack1) > 0 && len(stack2) > 0 || p != nil && q != nil {
		for p != nil && q != nil {
			stack1 = append(stack1, p)
			stack2 = append(stack2, q)
			p = p.Left
			q = q.Right
		}

		if p == nil && q != nil || p != nil && q == nil {
			return false
		}

		p = stack1[len(stack1)-1]
		q = stack2[len(stack2)-1]
		stack1 = stack1[:len(stack1)-1]
		stack2 = stack2[:len(stack2)-1]

		if p.Val != q.Val {
			return false
		}

		p = p.Right
		q = q.Left
	}

	if len(stack1) == 0 && len(stack2) == 0 {
		return true
	} else {
		return false
	}
}



