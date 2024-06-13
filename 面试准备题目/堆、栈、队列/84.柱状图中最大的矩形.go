package main

import (
	"fmt"
	"math"
)

// 暴力法求解(超出时间限制)
func largestRectangleArea1(heights []int) int {
	n := len(heights)
	if n == 1 {
		return (1 * heights[0])
	}
	max := 0
	for i := 0; i < n-1; i++{
		for j := i+1; j < n; j ++{
			x := j - i + 1
			y := math.MaxInt64
			tmp := 0
			for k := i; k <= j; k++{
				if heights[k] > tmp {
					tmp = heights[k]
				}
				if heights[k] < y {
					y = heights[k]
				}
			}
			if tmp > max {
				max = tmp
			}
			if x * y > max {
				max = x * y
			}
		}
	}
	return max
}

// 优化的暴力

/*
使用单调栈：

单调栈是一种特殊的栈，它保证栈中的元素始终是单调递增或单调递减的。
在这个问题中，我们使用单调递增栈来帮助找到每个柱子的左边界和右边界。
定义两个辅助数组：

left[i]：表示以 height[i] 为高的矩形在左侧延伸的最远位置。
right[i]：表示以 height[i] 为高的矩形在右侧延伸的最远位置。
初始化 right 数组为 n，表示默认右边界为数组的右边界。
遍历数组，填充 left 和 right 数组：

使用单调栈从左到右遍历数组，计算每个柱子的左边界和右边界。
对于每个柱子，更新其右边界，并记录其左边界。
计算最大矩形面积：

使用 left 和 right 数组计算每个柱子的最大宽度，进而计算最大矩形面积。

从左到右遍历 height 数组，使用单调栈来计算每个柱子的左边界和右边界。
在遍历过程中，如果当前柱子高度小于或等于栈顶元素的高度，则更新栈顶元素的右边界为当前柱子的下标，并将栈顶元素弹出。
计算当前柱子的左边界。如果栈为空，左边界设为 -1；否则，左边界设为栈顶元素的下标。
将当前柱子的下标压入栈中。
 */
// 栈的解决方法
func largestRectangleArea(height []int) int{
	n := len(height)
	left, right := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		right[i] = n
	}
	mono_stack := []int{}
	for i := 0; i < n; i++ {
		for len(mono_stack) > 0 && height[mono_stack[len(mono_stack)-1]] >= height[i] {
			right[mono_stack[len(mono_stack)-1]] = i
			mono_stack = mono_stack[:len(mono_stack)-1]
		}
		if len(mono_stack) == 0{
			left[i] = -1
		}else{
			left[i] = mono_stack[len(mono_stack)-1]
		}
		mono_stack = append(mono_stack, i)
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans = max(ans, (right[i]-left[i]-1)*height[i])
	}
	return ans
}

func largestRectangleArea(heights []int) int {
	res := 0
	for i := 0; i < len(heights); i++ {
		tmp := getArea(heights, i)
		if  tmp > res {
			res = tmp
		}
	}
	return res
}

func getArea(heights []int, index int) int {
	left := index - 1
	right := index + 1
	count := 1

	for left >= 0 {
		if heights[left] >= heights[index] {
			count++
			left--
		} else {
			break
		}
	}

	for right < len(heights) {
		if heights[right] >= heights[index] {
			count++
			right++
		} else {
			break
		}
	}

	return count * heights[index]
}

func max(x, y int) int{
	if x > y{
		return x
	}else{
		return y
	}
}

func main(){
	s := []int{1,4,7,2,9,3,8}
	out := largestRectangleArea(s)
	fmt.Println(out)
}