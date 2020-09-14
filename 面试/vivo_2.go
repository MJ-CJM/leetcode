package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	_, _ = fmt.Scan(&s)
	res := process_huiwen(s)
	fmt.Printf("%s", res)
}

func process_huiwen(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}
	for i := 0; i < n; i++ {
		tmp := s[0:i] + s[i+1:]
		if isPalindrome_1(tmp) {
			return tmp
		}
	}
	return "false"
}

func isPalindrome_1(s string) bool {
	s = strings.ToLower(s)
	n := len(s)
	i := 0
	j := n - 1
	for  i < j {
		if ((s[i] >= 'a' && s[i] <= 'z') || (s[i] >= '0' && s[i] <= '9')) &&
			((s[j] >= 'a' && s[j] <= 'z') || (s[j] >= '0' && s[j] <= '9')){
			if s[i] != s[j] {
				return false
			}else{
				i++
				j--
				continue
			}
		}
		if !((s[i] >= 'a' && s[i] <= 'z') || (s[i] >= '0' && s[i] <= '9')) {
			i++
		}
		if !((s[j] >= 'a' && s[j] <= 'z') || (s[j] >= '0' && s[j] <= '9')){
			j--
		}
	}
	return true
}
