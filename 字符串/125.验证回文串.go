package main

import "strings"

func isPalindrome(s string) bool {
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