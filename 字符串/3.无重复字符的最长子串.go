package main

import (
	"fmt"
	"reflect"
)

func lengthOfLongestSubstring(s string) int {
	n := len(s)
	res := 0
	start, end := 0, -1
	freq := make([]int, 128)
	for start < n {
		if end+1 < n && freq[s[end+1]] == 0 {
			end++
			freq[s[end]]++
		}else{
			freq[s[start]]--
			start++
		}
		res = max(res, end-start+1)
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