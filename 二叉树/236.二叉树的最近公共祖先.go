package main

// 递归
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 找不到（nil），或找到了（左子节点或右子节点）
	if root == nil || root.Val == p.Val || root.Val == q.Val{
		return root
	}
	// 左子树找
	l := lowestCommonAncestor(root.Left, p,q)

	// 右子树找
	r := lowestCommonAncestor(root.Right, p,q)

	if l == nil{
		return r
	}else if r == nil{
		return l
	}else{
		// 左右都找到了，那说明我就是公共祖先喽
		return root
	}
}