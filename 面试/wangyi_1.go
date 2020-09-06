package main

import (
	"fmt"
	"sort"
)

func main() {
	var x int
	var input []int
	_, _ = fmt.Scan(&x)
	for x > 0{
		input = append(input, x)
		_, _ = fmt.Scan(&x)
	}
	sort.Ints(input)
	output := process(input)
	fmt.Printf("%d", output)
}

func process(a []int) int {
	if a[0] >= 2 {
		return a[0] -1
	}
	for i := 1; i < len(a); i++ {
		if a[i] - a[i-1] != 1 {
			return a[i-1]+1
		}
	}
	return -1
}