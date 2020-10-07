package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	tmp := []int{}
	for i := 0; i < n; i++ {
		var x int
		_, _ = fmt.Scan(&x)
		tmp = append(tmp, x)
	}
	res := process_min_4(tmp)
	fmt.Printf("%d", res)
}

func process_min_4(tmp []int) int {
	sort.Ints(tmp)
	return tmp[3]
}
