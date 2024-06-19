package main

import (
	"fmt"
	"reflect"
)

func lengthOfLongestSubstring(s string) int {
	n := len(s)

	if n <= 1 {
		return n
	}

	res := 0
	sMap := make(map[byte]int)
	i := 0

	for j := 0; j < n; j++ {
		if v, ok := sMap[s[j]]; ok && v >= i {
			i = v + 1
		}

		sMap[s[j]] = j

		tmp := j - i + 1
		if tmp > res {
			res = tmp
		}
	}

	return res
}

func max(i,j int)int{
	if i>j{
		return i
	}else{
		return j
	}
}

func main() {
	fmt.Println(reflect.TypeOf('a'))
}