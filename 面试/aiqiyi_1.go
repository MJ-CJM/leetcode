package main

import "fmt"

func main() {
	var s string
	_, _ = fmt.Scan(&s)
	res := process_chongfu(s)
	fmt.Printf("%d", res)
}

func process_chongfu(s string) int {
	n := len(s)
	if n <= 1 {
		return n
	}
	res := 0
	for i := 0; i < n - 1; i++ {
		for j := i + 1; j < n; j++ {
			tmp := s[i:j]
			tmp2 := count_str(tmp)
			if tmp2 > 0 && tmp2 > res {
				res = tmp2
			}
		}
	}
	return res
}

func count_str(tmp string) int {
	s_l := make([]int, 26)
	n := len(tmp)
	for i := 0; i < n; i++ {
		if s_l[tmp[i] - 'a'] == 0 {
			s_l[tmp[i] - 'a'] = 1
		}else{
			return -1
		}
	}
	return n
}
