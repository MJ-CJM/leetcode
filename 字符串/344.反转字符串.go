package main

func reverseString(s []byte) {
	n := len(s)
	for n <= 1 {
		return
	}
	for p, q := 0, n-1; p < q; p, q = p+1, q-1 {
		s[p], s[q] = s[q], s[p]
	}
	return
}
