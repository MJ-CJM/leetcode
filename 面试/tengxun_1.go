package main

import "fmt"

func main () {
	n := 0
	l1 := []int{}
	_, _ = fmt.Scan(&n)
	for i := 0; i < n; i++ {
		x := 0
		_, _ = fmt.Scan(&x)
		l1 = append(l1, x)
	}
	m := 0
	l2 := []int{}
	_, _ = fmt.Scan(&m)
	for i := 0; i < m; i++ {
		x := 0
		_, _ = fmt.Scan(&x)
		l2 = append(l2, x)
	}
	out := process_lianbiao(l1, l2)
	for i := 0; i < len(out); i++ {
		fmt.Printf("%d ", out[i])
	}
}

func process_lianbiao(l1 []int, l2 []int) []int {
	n := len(l1)
	m := len(l2)
	min := 0
	max := 0
	if n <= m {
		min = n
		max = m
	}else{
		min = m
		max = n
		l1, l2 = l2, l1
	}
	res := []int{}
	tmp := 0
	for i := 0; i < min; i++ {
		if l1[i] > l2[tmp] {
			continue
		}
		for j := tmp; j < max; j++ {
			if l1[i] == l2[j] {
				res = append(res, l1[i])
				tmp = j+1
				break
			}
		}
	}
	return res
}
