package main

func pathSum(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 构建哈希表表示节点
	nodes := make(map[[2]int]int)
	for _, num := range nums {
		d := num / 100
		p := (num / 10) % 10
		v := num % 10
		nodes[[2]int{d, p}] = v
	}

	var dfs func(d, p int, currSum int) int
	dfs = func(d, p int, currSum int) int {
		if _, exists := nodes[[2]int{d, p}]; !exists {
			return 0
		}

		// 当前节点值
		currSum += nodes[[2]int{d, p}]

		// 计算左子节点和右子节点
		left := dfs(d+1, p*2-1, currSum)
		right := dfs(d+1, p*2, currSum)

		// 如果是叶子节点
		if left == 0 && right == 0 {
			return currSum
		}
		return left + right
	}

	// 从根节点开始DFS
	return dfs(1, 1, 0)
}
