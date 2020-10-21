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
