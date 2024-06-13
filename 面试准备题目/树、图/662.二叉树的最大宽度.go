package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Wnode struct {
	Node     *TreeNode
	Index    int
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []Wnode{Wnode{
		Node: root,
		Index: 1,
	}}

	res := 0

	for len(queue) != 0 {
		first := queue[0]
		end := queue[len(queue) - 1]
		tmp := end.Index - first.Index + 1
		if tmp > res {
			res = tmp
		}

		tmpQ := []Wnode{}

		for i := 0; i < len(queue); i++ {
			if queue[i].Node.Left != nil {
				tmpQ = append(tmpQ, Wnode{
					Node: queue[i].Node.Left,
					Index: queue[i].Index * 2,
				})
			}
			if queue[i].Node.Right != nil {
				tmpQ = append(tmpQ, Wnode{
					Node: queue[i].Node.Right,
					Index: queue[i].Index * 2 + 1,
				})
			}
		}
		queue = tmpQ
	}

	return res
}
