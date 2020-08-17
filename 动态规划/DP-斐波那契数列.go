package main

import "fmt"

// 递归
func fib(N int) int {
	if N <= 1 {
		return N
	} else {
		return fib(N-1) + fib(N-2)
	}
}

// 砍掉重复节点-使用缓存（自顶向下）
func fib2(N int) int {
	mem := make([]int, N+1)
	return fib_dp1(N, mem)
}

func fib_dp1(N int, mem []int) int {
	if N <= 1 {
		return N
	}
	if mem[N] == 0 {
		mem[N] = fib_dp1(N-1, mem) + fib_dp1(N-2, mem)
	}
	return mem[N]
}

//自底向上递推（动态规划的终极状态）
func fib_2(N int) int {
	if N <= 1 {
		return N
	}
	result := make([]int, N+1)
	result[0] = 0
	result[1] = 1
	for i := 2; i <= N; i++ {
		result[i] = result[i-1] + result[i-2]
	}
	return result[N]
}

func main() {
	n := 5
	fmt.Println(fib(n))
}
