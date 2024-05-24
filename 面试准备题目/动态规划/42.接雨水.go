package main

func trap(height []int) int {
	var left, right, leftMax, rightMax, res int
	right = len(height) - 1
	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				res += leftMax - height[left]
			}
			left++
		} else {
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				res += rightMax - height[right]
			}
			right--
		}
	}
	return res
}

func trap2(height []int) int {
	if len(height) == 0 {
		return 0
	}

	n := len(height)
	leftMax := make([]int, n)
	rightMax := make([]int, n)
	water := 0

	// 填充 leftMax 数组
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	// 填充 rightMax 数组
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	// 计算总的接雨水量
	for i := 0; i < n; i++ {
		water += min(leftMax[i], rightMax[i]) - height[i]
	}

	return water
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
