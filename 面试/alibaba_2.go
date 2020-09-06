package main

import "fmt"

func main() {
	var a, b, A, B int
	_, _ = fmt.Scan(&A, &B, &a, &b)
	x, y := max_compute(A, B, a, b)
	fmt.Printf("%d %d", x, y)
}

func max_compute(A int, B int, a int, b int) (int, int) {
	var x int
	var y int
	for i := B; i >= 1; i-- {
		if i * a % b == 0 && i * a / b >= 1 && i * a / b <= A{
			x = i * a / b
			y = i
			break
		}else{
			i = i - i * a % b - 1
		}

	}

	return x, y
}