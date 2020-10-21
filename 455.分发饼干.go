package main

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	n1 := len(g)
	n2 := len(s)
	c1 := 0
	c2 := 0
	count := 0
	flag := true
	for flag {
		if c1 == n1 || c2 == n2{
			break
		}
		if g[c1] <= s[c2]{
			c1 ++
			c2 ++
			count ++
		}else{
			c2 ++
		}
	}
	return count
}


