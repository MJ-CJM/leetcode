package main

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	// 双端队列，存储数组的索引
	deque := []int{}
	result := []int{}

	for i := 0; i < len(nums); i++ {
		// 移除不在窗口范围内的元素
		if len(deque) > 0 && deque[0] < i-k+1 {
			deque = deque[1:]
		}

		// 移除所有小于当前元素的元素
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}

		// 添加当前元素的索引
		deque = append(deque, i)

		// 当前窗口的最大值是队列头部的元素
		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}

	return result
}

