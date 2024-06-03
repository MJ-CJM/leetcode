package main

import "fmt"

func firstUniqChar(s string) int {
	mp := make(map[int32]int)
	for _, v := range s {
		_, ok := mp[v]
		if ok {
			mp[v]++
		} else {
			mp[v] = 1
		}
	}
	for i, v := range s {
		if mp[v] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	s := "leetcode"
	fmt.Println(firstUniqChar(s))
}
