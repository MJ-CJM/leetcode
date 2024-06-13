package main

func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}
	j := 0

	for _, val := range pushed {
		stack = append(stack, val) // 将元素压入栈
		// 检查栈顶元素是否与 popped 序列的当前元素相同
		for len(stack) > 0 && stack[len(stack)-1] == popped[j] {
			stack = stack[:len(stack)-1] // 弹出栈顶元素
			j++ // 移动 popped 序列的指针
		}
	}

	// 如果栈为空且 popped 序列的所有元素都已经匹配，则返回 true
	return len(stack) == 0
}
