package main

import (
	"fmt"
)

func main() {
	var c byte
	var err error
	var b []int
	for ; err == nil; {
		_, err = fmt.Scanf("%c", &c)

		if c != '\n' {
			if c != ' ' {
				b = append(b, int(c - '0'))
			}
		} else {
			break;
		}
	}
	out := process_2(b)
	fmt.Printf("%d", out)
}


func process_2(nums []int) int {
	res := []int{}
	res = append(res, 0)
	for _, v := range nums {
		res = append(res, v)
	}
	n := len(res)
	for i := 1; i < n; i++ {
		l := 2 * i
		r := 2 * i + 1
		if l < n {
			if res[i] <= res[l] {
				return 0
			}
		}
		if r < n {
			if res[i] >= res[r] {
				return 0
			}
		}
	}
	return 1
}
