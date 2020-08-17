package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	x, list := process_Nixu(n)
	if x == 0 {
		fmt.Println("0")
		return
	}
	fmt.Println(x)
	for i := 0; i < x; i++ {
		fmt.Printf("%d %d\n", list[0], list[1])
		list = list[2:]
	}
}

func process_Nixu(n int) (int, []int) {
	count := 0
	list := []int{}
	for i := 1; i <= n; i++ {
		if 4*i > n {
			break
		}
		if is_nixu(i, 4*i) {
			count++
			list = append(list, i)
			list = append(list, 4*i)
		}
	}
	return count, list
}

func is_nixu(x int, y int) bool {
	s1 := strconv.Itoa(x)
	s2 := strconv.Itoa(y)
	n1 := len(s1)
	n2 := len(s2)
	if n1 != n2 {
		return false
	}
	for i := 0; i < n1; i++ {
		if s1[i] != s2[n1-i-1] {
			return false
		}
	}
	return true
}
