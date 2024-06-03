package main

func validPalindrome(s string) bool {
	i, j := 0, len(s)-1
	i, j, f := isValid(s, i, j)
	if !f {
		_, _, f1 := isValid(s, i+1, j)
		_, _, f2 := isValid(s, i, j-1)
		return f1 || f2
	}
	return true
}

func isValid(s string, i int, j int) (int, int, bool) {
	for i < j {
		if s[i] != s[j] {
			return i, j, false
		}else{
			i++
			j--
		}
	}
	return 0, 0, true
}