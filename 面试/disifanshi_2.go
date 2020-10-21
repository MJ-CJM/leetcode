package main

import (
	"fmt"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	tmp := []int{}
	max := 0
	for i := 0; i < n; i++ {
		var x,y int
		_, _ = fmt.Scan(&x, &y)
		if x > max {
			max = x
		}
		if y > max {
			max = y
		}
		tmp = append(tmp, x)
		tmp = append(tmp, y)
	}
	res := process_fugai(tmp, max)
	fmt.Printf("%d", res)
}

func process_fugai(tmp []int, max int) int {
	example := make([]int, max+1)
	n := len(tmp)
	for i := 0; i < n; i++ {
		if example[tmp[i]] == 0 {
			example[tmp[i]] = 1
		}
	}
	max_int := 0
	j := 0
	for j < max {
		tmp2 := 1
		for j < max && j+1 < max && example[j] == 1 && example[j+1] == 1 {
			tmp2++
			j++
		}
		if tmp2 == 1 {
			j++
			continue
		}
		if tmp2 > max_int {
			max_int = tmp2
		}
	}
	return max_int
}

//package main
//
//import "fmt"
//
//func main() {
//	var n int
//	_, _ = fmt.Scan(&n)
//	tmp := [][]int{}
//	for i := 0; i < n; i++ {
//		var x,y int
//		_, _ = fmt.Scan(&x, &y)
//		tmp = append(tmp, []int{x, y})
//	}
//	res := process_fugai(tmp)
//	fmt.Printf("%d", res)
//}
//
//func process_fugai(tmp [][]int) int {
//	max := 0
//	n := len(tmp)
//	i := 0
//	for i < n-1 {
//		tmp2 := 1
//		for i < n && i+1 < n && tmp[i][1] == tmp[i+1][0] {
//			tmp2++
//			i++
//		}
//		if tmp2 == 1 {
//			i++
//			continue
//		}
//		if tmp2 > max {
//			max = tmp2
//		}
//	}
//	return max
//}

