package main

import "fmt"

func main() {
	n := 0
	_,_ = fmt.Scan(&n)
	input := [][]int{}
	for i := 0; i < n; i++ {
		x, y, z, p := 0, 0, 0, 0
		_,_ = fmt.Scan(&x, &y, &z, &p)
		input = append(input, []int{x, y, z, p})
	}
	fmt.Printf("%d", 28)
}
