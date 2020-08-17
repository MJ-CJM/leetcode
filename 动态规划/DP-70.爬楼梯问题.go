package main

func climbStairs(n int) int {
	if n <= 2{
		return 2
	}
	f1 := 1
	f2 := 2
	f3 := 3
	for i := 3; i <= n; i++ {
		f3 = f1 + f2
		f1 = f2
		f2 = f3
	}
	return f3
}