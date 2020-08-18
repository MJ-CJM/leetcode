package main

import "strings"

func reverseWords3(s string) string {
	str := []byte(s)
	reverseString2(str)
	s = string(str)
	ls := strings.Split(s, " ")
	n := len(ls)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		ls[i], ls[j] = ls[j], ls[i]
	}
	return strings.Join(ls, " ")
}

func reverseString2(s []byte) {
	n := len(s)
	for n <= 1 {
		return
	}
	for p, q := 0, n-1; p < q; p, q = p+1, q-1 {
		s[p], s[q] = s[q], s[p]
	}
	return
}