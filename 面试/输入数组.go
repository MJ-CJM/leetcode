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
	fmt.Println(b)
}