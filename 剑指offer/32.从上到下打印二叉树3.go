package main

func levelOrder_3(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := []*TreeNode{root}
	res := [][]int{}
	flag := 0
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
		if flag == 1 {
			Reverse(tmp)
			flag = 0
		}else{
			flag = 1
		}
		res = append(res, tmp)
		queue = tmp2
	}
	return res
}

func Reverse(nums []int) {
	n := len(nums)
	if n == 1 {
		return
	}
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return
}
