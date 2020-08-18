package main

func reverseOnlyLetters(S string) string {
	buf := []byte(S)
	for i, j := 0, len(buf)-1; i < j; {
		for i < j && !isLetter(buf[i]) {
			i++
		}
		for i < j && !isLetter(buf[j]) {
			j--
		}
		buf[i], buf[j] = buf[j], buf[i]
		i++
		j--
	}
	return string(buf)
}

func isLetter(c byte) bool {
	if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' {
		return true
	}
	return false
}