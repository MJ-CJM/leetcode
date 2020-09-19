package main

import "fmt"

func main(){
	a := 'A'
	fmt.Println(string(a+1))
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