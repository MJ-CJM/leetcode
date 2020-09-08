package main

import "fmt"

func main() {
	n := 0
	_, _ = fmt.Scan(&n)
	s := ""
	_, _ = fmt.Scan(&s)
	num := []int{}
	for i := 0; i < 2 * n; i++ {
		x := 0
		_, _ = fmt.Scan(&x)
		if s[i] == 'B' {
			x += 3000
		}
		num = append(num, x)
	}
	res := process_caozuo(num)
	fmt.Printf("%d", res)
}

func process_caozuo(num []int) int {
	return 0
}



