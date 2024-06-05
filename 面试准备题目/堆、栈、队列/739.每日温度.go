package main

// 单调栈逻辑处理
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	stack := []int{}
	res := make([]int, n)
	for i := 0; i < n; i++ {
		cur := temperatures[i]
		for len(stack) > 0 && cur > temperatures[stack[len(stack) - 1]] {
			pre := stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			res[pre] = i - pre
		}
		stack = append(stack, i)
	}
	return res
}
