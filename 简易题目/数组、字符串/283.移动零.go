package main

import "fmt"

func moveZeroes(nums []int)  {
	n := len(nums)
	j := 0
	for i := 0; i < n; i++ {
		if nums[i] != 0{
			tmp := nums[j]
			nums[j] = nums[i]
			nums[i] = tmp
			j++
		}
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}