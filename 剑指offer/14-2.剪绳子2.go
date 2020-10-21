package main

func cuttingRope_2(n int) int {
	if n <= 3 {
		return n-1
	}
	ret := 1
	for n > 4 {
		ret = ret * 3 % 1000000007
		n -= 3
	}
	return ret * n % 1000000007
}
