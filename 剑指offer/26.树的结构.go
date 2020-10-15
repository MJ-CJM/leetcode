package main

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	if A == nil || B == nil {
		return false
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, A)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.Val == B.Val {
			if is_matchtree(node, B) {
				return true
			}
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return false
}

func is_matchtree(a *TreeNode, b *TreeNode) bool {
	if b == nil {
		return true
	}
	if a == nil {
		return false
	}
	if a.Val != b.Val {
		return false
	}
	return is_matchtree(a.Left, b.Left) && is_matchtree(a.Right, b.Right)
}


