package main

import "fmt"

func main(){
	l := []int{1, 2, 2, 3}
	l = quchong1(l)
	fmt.Println(l)
}

func quchong1(list []int) []int {
	res := []int{}
	for i := 0; i < len(list); i++ {
		if (i > 0 && list[i-1] == list[i]) {
			continue
		}
		res = append(res, list[i])
	}
	return res
}