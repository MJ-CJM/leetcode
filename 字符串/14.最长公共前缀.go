package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	n := len(strs)
	if n == 0 {
		return ""
	}
	if n == 1 {
		return strs[0]
	}
	tmp := strs[0]
	m := len(tmp)
	count := 0
	for i := 0; i < m; i++ {
		flag := 0
		for j := 1; j < n; j++ {
			if i >= len(strs[j]) {
				flag = 1
				break
			}
			if strs[j][i] != tmp[i] {
				flag = 1
				break
			}
		}
		if flag == 1 {
			break
		}
		count++
	}
	if count == 0 {
		return ""
	}
	res := ""
	for i := 0; i < count; i++ {
		res += string(tmp[i])
	}
	return res
}

func longest_2(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	prefix := strs[0]
	for _, k := range strs {
		for strings.Index(k, prefix) != 0 {
			if len(prefix) == 0{
				return ""
			}
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}

func main() {
	str := []string{"flower","flow","flight"}
	fmt.Println(longestCommonPrefix(str))
}
