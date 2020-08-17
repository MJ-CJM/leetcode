package main

import (
	"fmt"
)

func lengthOfLastWord(s string) int {
	if s == "" {
		return 0
	}
	n := len(s)
	res := 0
	for i := n-1; i >= 0; i-- {
		if string(s[i]) != " " {
			res++
		}
		if res != 0 && string(s[i]) == " "{
			break
		}
	}
	return res
}

func main() {
	s := "a "
	fmt.Println(lengthOfLastWord(s))
}