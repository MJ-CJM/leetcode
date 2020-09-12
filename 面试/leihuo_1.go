package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	_, _ = fmt.Scan(&s)
	res := process_leihuo(s)
	fmt.Printf("%d", res)
}

func process_leihuo(s string) int {
	n := len(s)
	if n <= 0 {
		return 0
	}
	count := 0
	for i := 0; i <= n-1; i++ {
		for j := i+1; j <= n; j++ {
			tmp := s[i:j]
			if isPalindrome(tmp) {
				count++
			}
		}
	}
	return count
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	n := len(s)
	if n <= 1 {
		return false
	}
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