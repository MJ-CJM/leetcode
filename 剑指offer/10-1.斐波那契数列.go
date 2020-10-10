package main

// 迭代推导
func fib(n int) int {
	if n <= 1 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

// 递归
func fib_2(n int) int {
	if n <= 1 {
		return n
	}else{
		return fib_2(n-1) + fib_2(n-2)
	}
}
