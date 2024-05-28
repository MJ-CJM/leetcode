package main

import (
	"fmt"
)

// 动态规划：
// 重复子问题：a[i] = max[i-1]
// 定义状态：a[i][0,1] 0: 不偷，1：偷
// 状态转移方程：a[i][0] = max(a[i-1][0], a[i-1][i])
//               a[i][1] = a[i-1][0] + nums[i]
func rob(nums []int) int {
	if nums == nil || len(nums) == 0{
		return 0
	}
	n := len(nums)
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, 2)
	}
	a[0][0] = 0
	a[0][1] = nums[0]
	for i := 1; i < n; i++ {
		a[i][0] = max_int2(a[i-1][0], a[i-1][1])
		a[i][1] = a[i-1][0] + nums[i]
	}
	return max_int2(a[n-1][0], a[n-1][1])
}

func rob2(nums []int) int{
	if nums == nil || len(nums) == 0{
		return 0
	}
	if len(nums) == 1{
		return nums[0]
	}
	n := len(nums)
	a := make([]int, n)
	a[0] = nums[0]
	a[1] = max_int2(nums[0], nums[1])
	for i := 2; i < n; i++ {
		a[i] = max_int2(a[i-1], a[i-2]+nums[i])
	}
	return a[n-1]
}

func max_int2(x int, y int)int{
	if x >= y{
		return x
	}else{
		return y
	}
}

func main(){
	nums := []int{1,2,3,1}
	fmt.Println(rob(nums))
}