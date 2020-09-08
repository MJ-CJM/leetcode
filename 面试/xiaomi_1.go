package main

import (
	"fmt"
	"strings"
)

func main() {
	var c byte
	var err error
	var b []string
	count_1 := 0
	var jian []int
	for ; err == nil; {
		_, err = fmt.Scanf("%c", &c)

		if c != '\n' {
			if c != ' ' {
				b = append(b, string(c))
				count_1++
			}else{
				b = append(b, string(c))
				jian = append(jian, count_1)
				count_1++
			}
		} else {
			break;
		}
	}
	input := []string{}
	tmp := 0
	for i := 0; i < len(jian); i++ {
		input = append(input, strings.Join(b[tmp:jian[i]], ""))
		tmp = jian[i]
	}
	input = append(input, strings.Join(b[tmp:], ""))
	result := panduanmima(input)
	for _, v := range result {
		fmt.Printf("%d \n", v)
	}
}

func panduanmima(input []string) []int {
	n := len(input)
	result := []int{}
	for i := 0; i < n; i++ {
		tmp := process_panduanmima(input[i])
		result = append(result, tmp)
	}
	return result
}

func process_panduanmima(s string) int {
	if len(s) < 8 || len(s) > 20 {
		return 1
	}
	flag := []int{0, 0, 0, 0}
	for _, v := range s {
		if v >= '0' && v <= '9' {
			if flag[0] != 1 {
				flag[0] = 1
			}
		}else if v >= 'A' && v <= 'Z' {
			if flag[2] != 1 {
				flag[2] = 1
			}
		}else if v >= 'a' && v <= 'z' {
			if flag[3] != 1 {
				flag[3] = 1
			}
		}else{
			if flag[1] != 1 {
				flag[1] = 1
			}
		}
		if flag[0] + flag[1] + flag[2] + flag[3] == 4 {
			return 0
		}
	}
	if flag[0] + flag[1] + flag[2] + flag[3] != 4 {
		return 2
	}
	return 0
}