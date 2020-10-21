package main

import (
	"fmt"
	"sort"
)

func main() {
	n := 0
	_, _ =fmt.Scan(&n)
	input := []int{}
	for i := 0; i < n; i++ {
		x := 0
		_, _ =fmt.Scan(&x)
		input = append(input, x)
	}
	output := process_sort(input)
	fmt.Printf("%d \n%d", output[0], output[1])
}

func process_sort(input []int) []int {
	n := len(input)
	if n <= 1{
		return []int{-1, -1}
	}
	if n == 2 {
		if input[0] <= input[1] {
			return []int{-1, -1}
		}else{
			return []int{0, 1}
		}
	}
	for i := 0; i < n-1; i++ {
		for j := i+1; j < n; j++ {
			if panduan(input[:i]) && panduan(input[j+1:]) && !panduan(input[i:j+1]){
				tmp1 := input[i:j+1]
				sort.Ints(tmp1)
				tmp2 := append(append(input[:i], tmp1...), input[j+1:]...)
				if panduan(tmp2) {
					flag := 0
					for k := 0; k < n; k++ {
						if !panduan(input[:k]) {
							flag = k-1
						}
					}
					return []int{flag, j}
				}
			}
		}
	}
	return []int{-1, -1}
}

func panduan(nums []int) bool {
	n := len(nums)
	if n <= 1 {
		return true
	}
	for i := 1; i < n; i++ {
		if nums[i] < nums[i-1] {
			return false
		}
	}
	return true
}
