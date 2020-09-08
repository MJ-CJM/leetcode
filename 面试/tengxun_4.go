package main

import (
	"fmt"
)

func main() {
	n := 0
	l := []int{}
	_, _ = fmt.Scan(&n)
	for i := 0; i < n; i++ {
		x := 0
		_, _ = fmt.Scan(&x)
		l = append(l, x)
	}
	output := process_zhongwei(l)
	for i := 0; i < n; i++ {
		fmt.Printf("%d \n", output[i])
	}
}

func process_zhongwei(nums []int) []int {
	res := []int{}
	n := len(nums)
	sort_nums := QuickSort(nums)
	l := sort_nums[n>>1 - 1]
	r := sort_nums[n>>1]
	for _, v := range nums {
		if v <= l {
			res = append(res, r)
		}else{
			res = append(res, l)
		}
	}
	return res
}

func QuickSort(a []int) []int {
	if len(a) <= 1 {
		return a
	} else {
		head := a[0]
		left, right := []int{}, []int{}
		for i := 1; i < len(a); i++ {
			if a[i] < head {
				left = append(left, a[i])
			} else {
				right = append(right, a[i])
			}
		}
		return append(append(QuickSort(left), head), QuickSort(right)...)
	}
}