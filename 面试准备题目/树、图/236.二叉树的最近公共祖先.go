package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 如果当前节点为null，返回null
	if root == nil {
		return nil
	}

	// 如果当前节点是p或q，返回当前节点
	if root == p || root == q {
		return root
	}

	// 在左子树中递归查找
	left := lowestCommonAncestor(root.Left, p, q)
	// 在右子树中递归查找
	right := lowestCommonAncestor(root.Right, p, q)

	// 如果左子树和右子树都找到了节点，则当前节点是最近公共祖先
	if left != nil && right != nil {
		return root
	}

	// 否则，返回非空的子节点
	if left != nil {
		return left
	}
	return right
}
