package main

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var result []int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if i == levelSize-1 {
				result = append(result, node.Val)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return result
}


func rightSideView2(root *TreeNode) []int {
	var result []int
	dfs(root, 0, &result)
	return result
}

func dfs(node *TreeNode, depth int, result *[]int) {
	if node == nil {
		return
	}

	if depth == len(*result) {
		*result = append(*result, node.Val)
	}

	dfs(node.Right, depth+1, result)
	dfs(node.Left, depth+1, result)
}