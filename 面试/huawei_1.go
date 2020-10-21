package main

import "fmt"

func main() {
	var result int
	n := 0
	_, _ = fmt.Scan(&n)
	if n < 1 || n > 100000 {
		result = -1
	}
	A := []int{}
	B := []int{}
	for i := 0; i < n; i++ {
		x := 0
		_, _ = fmt.Scan(&x)
		if x < 0 || x > 5 {
			result = -1
		}
		A = append(A, x)
	}
	for i := 0; i < n; i++ {
		x := 0
		_, _ = fmt.Scan(&x)
		if x < 0 || x > 5 {
			result = -1
		}
		B = append(B, x)
	}
	m := 0
	_, _ = fmt.Scan(&m)
	if m < 1 || m > 1000000 {
		result = -1
	}
	C := []int{}
	D := []int{}
	for i := 0; i < m; i++ {
		x := 0
		_, _ = fmt.Scan(&x)
		if x < 0 || x > 5 {
			result = -1
		}
		C = append(C, x)
	}
	for i := 0; i < m; i++ {
		x := 0
		_, _ = fmt.Scan(&x)
		if x < 0 || x > 5 {
			result = -1
		}
		D = append(D, x)
	}
	if result != -1 {
		result = process_perfect(A, B, C, D)
	}
	fmt.Printf("%d", result+1)
}

func process_perfect(a []int, b []int, c []int, d []int) int {
	n := len(a)
	m := len(c)
	hashmap := make(map[int]int)
	for i := 0; i < n; i++ {
		hashmap[a[i]] = b[i]
	}
	count := 0
	for i := 0; i < m; i++ {
		if hashmap[c[i]] == d[i] {
			count++
			if count == n {
				return i - n + 1
			}
		}else{
			count = 0
		}
	}
	return -1
}


//import sys
//def print1(p, a, b, q, c, d):
//for i in range (q-p):
//if c[i:i+p] == a and d[i:i+p] == b:
//return i + 1
//return 0
//
//line1 = sys.stdin.readline().strip()
//p = int(line1)
//line2 = sys.stdin.readline().strip()
//values2 = list(line2.split())
//a = "".join(values2)
//line3 = sys.stdin.readline().strip()
//values3 = list(line3.split())
//b = "".join(values3)
//line4 = sys.stdin.readline().strip()
//q = int(line4)
//line5 = sys.stdin.readline().strip()
//values5 = list(line5.split())
//c = "".join(values5)
//line6 = sys.stdin.readline().strip()
//values6 = list(line6.split())
//d = "".join(values6)
//
//print(print1(p, a, b, q, c, d))
