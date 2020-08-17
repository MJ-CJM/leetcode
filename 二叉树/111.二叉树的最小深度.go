package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil{
		return 0
	}
	count := 0
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		count++
		temp := []*TreeNode{}
		for _, v := range queue {
			if v.Left == nil && v.Right == nil{
				return count
			}
			if v.Left != nil{
				temp = append(temp, v.Left)
			}
			if v.Right != nil{
				temp = append(temp, v.Right)
			}
		}
		queue = temp
	}
	return count
}
