package main

import (
	"fmt"
	"strings"
)

func main() {
	n := 0
	_, _ = fmt.Scan(&n)
	//s := ""
	//_, _ = fmt.Scan(&s)
	var c byte
	var err error
	var d []string
	for ; err == nil; {
		_, err = fmt.Scanf("%c", &c)

		str := string(c)
		if str != "\n" {
			d = append(d, str)
		} else {
			break;
		}
	}
	s := strings.Join(d, "")
	res := process_pojie(n, s)
	fmt.Printf("%s", res)
}

func process_pojie(n int, s string) string {
	m := len(s)
	if m == 0{
		return ""
	}
	res := ""
	for i := 0; i < m; i++ {
		if (i+1) % n == 0 {
			tmp := s[i-n+1:i+1]
			tmp2 := reverse_string(tmp)
			res += tmp2
			continue
		}
		if (i+1)/n == m/n {
			tmp := s[i:]
			tmp2 := reverse_string(tmp)
			res += tmp2
			break
		}
	}
	return res
}

func reverse_string(s string) string {
	n := len(s)
	res := ""
	for i := n-1; i >= 0; i-- {
		res += string(s[i])
	}
	return res
}
