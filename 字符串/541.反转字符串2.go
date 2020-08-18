package main

func reverseStr(s string, k int) string {
	str := []byte(s)
	i := 0
	for i < len(s) {
		l := i
		r := i+k-1
		if r >= len(str) - 1 {
			r = len(str)-1
		}
		for l < r && l < len(str) {
			str[l], str[r] = str[r], str[l]
			l++
			r--
		}
		i += 2 * k
	}
	return string(str)
}
