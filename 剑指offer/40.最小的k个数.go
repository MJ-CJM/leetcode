package main

import (
	"fmt"
	"sort"
)

func getLeastNumbers(arr []int, k int) []int {
	sort.Ints(arr)
	return arr[:k]
}

func main() {
	arr := []int{3, 2, 1}
	k := 2
	fmt.Println(getLeastNumbers(arr, k))
}
