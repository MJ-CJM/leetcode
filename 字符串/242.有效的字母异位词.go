package main

func isAnagram(s string, t string) bool {
	l1 := [26]int{}
	l2 := [26]int{}
	for i := 0; i < len(s); i++ {
		l1[(s[i]-'a')]++
	}
	for i := 0; i < len(t); i++ {
		l2[(t[i]-'a')]++
	}
	if l1 == l2 {
		return true
	}
	return false
}