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