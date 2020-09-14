package main

import (
	"fmt"
	"strconv"
	"strings"
)

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * 编译顺序
 * @param input string字符串
 * @return string字符串
 */
func compileSeq( input string ) string {
	// write code here
	str := strings.Split(input, ",")
	n := len(str)
	if n == 1 {
		return "0"
	}
	output := []string{}
	target := []string{"-1"}
	for i := 0; i < n; i++ {
		j_str := []string{}
		for j := 0; j < n; j++ {
			if is_in_seq(str[j], target) {
				j_s := strconv.Itoa(j)
				output = append(output, j_s)
				j_str = append(j_str, j_s)
				str[j] = "-2"
			}
		}
		target = j_str
	}
	return strings.Join(output, ",")
}

func is_in_seq(s string, target []string) bool {
	for _, v := range target {
		if v == s {
			return true
		}
	}
	return false
}

func main() {
	s := "1,2,-1,1"
	fmt.Println(compileSeq(s))
}
