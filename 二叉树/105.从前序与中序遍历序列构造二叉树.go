package main

// 根据前序数组找父节点，中序数组找出父节点对应的左右子树的思想来解决问题
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	// 用preorder找根节点
	root := &TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}
	i := 0
	// 找中序的根节点及其左右子树
	for ; i < len(inorder); i++ {
		if preorder[0] == inorder[i] {
			break
		}
	}
	// 找左子树
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	// 找右子树
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}
