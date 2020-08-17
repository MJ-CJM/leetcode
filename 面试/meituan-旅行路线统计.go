package main

import (
	"fmt"
)

func main() {
	var n int
	_, _ = fmt.Scan(&n)
	var list []string
	for i := 0; i < n; i++ {
		var s1, s2 string
		_, _ = fmt.Scan(&s1, &s2)
		list = append(list, []string{s1, s2}...)
	}
	count := count_luxian(list)
	fmt.Println(count)
}

func count_luxian(list []string) int {
	n := len(list)
	if n == 0{
		return 0
	}
	count := 0
	tmp := list[0]
	for i := 1; i < n; i++ {
		if list[i] == tmp {
			count++
			if i+1 < n{
				tmp = list[i+1]
				i++
			}else{
				break
			}
		}
	}
	return count
}
