package main

import "fmt"

func main() {
	n := 0
	k := 0
	_, _ = fmt.Scan(&n, &k)
	l := []string{}
	for i := 0; i < n; i++ {
		x := ""
		_, _ = fmt.Scan(&x)
		l = append(l, x)
	}

}

func compare(s1 string, s2 string) {

}