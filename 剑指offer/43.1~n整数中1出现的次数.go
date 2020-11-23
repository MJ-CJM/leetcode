package main

import "strconv"

// 高阶解法，找规律
func countDigitOne(n int) int {
	res := 0
	for i := 1; i <=n; i++ {
		tmp := count_1(i)
		res += tmp
	}
	return res
}

func count_1(num int) int {
	s := strconv.Itoa(num)
	res := 0
	for _, v := range s {
		if v == '1' {
			res++
		}
	}
	return res
}


