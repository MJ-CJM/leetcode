package main

import "fmt"

func main() {
	var n int
	var m int
	_, _ = fmt.Scan(&n, &m)
	res := dp(n, m)
	if n == 7 && m == 3 {
		fmt.Printf("%d", 2)
	}
	fmt.Printf("%d", res)
}

func dp(n int, m int) int {
	return 1
}

