package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	str := strings.Split(s, " ")
	tmp := 0
	for tmp < len(str)-1 {
		if str[tmp] == ""{
			str = append(str[:tmp], str[tmp+1:]...)
		}else{
			tmp++
		}
	}
	n := len(str)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		str[i], str[j] = str[j], str[i]
	}
	res := strings.Join(str, " ")
	return res
}

func main() {
	s := "a good   example"
	fmt.Println(reverseWords(s))
}
