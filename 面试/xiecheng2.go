package main

import (
	"fmt"
	"strings"
)

func main() {
	var c byte
	var err error
	var b []string
	for ; err == nil; {
		_, err = fmt.Scanf("%c", &c)

		if c != '\n' {
			b = append(b, string(c))
		} else {
			break;
		}
	}
	input := []string{}
	tmp := 0
	for i := 0; i < len(b); i++ {
		if b[i] == " " {
			input = append(input, strings.Join(b[tmp:i], ""))
			tmp = i + 1
		}
	}
	input = append(input, strings.Join(b[tmp:], ""))
	result := process_xunhuanyilai(input)
	for _, v := range result {
		fmt.Printf("%s \n", v)
	}
}

func process_xunhuanyilai(input []string) []string {
	n := len(input)
	level := 0
	iterm := ""
	result := []string{}
	dfs_xunhuan(input, n, level, iterm, &result)
	return result
}

func dfs_xunhuan(input []string, n int, level int, iterm string, result *[]string) {
	// terminator
	if level >= n {
		if !panduan_iterm(iterm) {
			tmp := iterm + "--circular dependency"
			*result = append(*result, tmp)
		}else{
			*result = append(*result, iterm)
		}
		return
	}
	// process && drill down
	value := input[level]
	for _, v := range value {
		c := iterm + string(v)
		dfs_xunhuan(input, n, level + 1, c, result)
	}
}

func panduan_iterm(iterm string) bool {
	s1 := make([]int, 26)
	for _, v := range iterm {
		if s1[v - 'a'] != 1 {
			s1[v - 'a'] = 1
		}else{
			return false
		}
	}
	return true
}