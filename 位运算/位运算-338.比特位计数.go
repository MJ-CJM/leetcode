package main

// 动态规划
// 重复子问题: res[i] = res[i-1] + nums[i]
// 定义状态: res[i]
// 状态转移方程 res[i] = res[i-1] + nums[i]
func countBits(num int) []int {
	res := []int{}
	for i := 0; i <= num; i++ {
		res = append(res, conunt_num(i))
	}
	return res
}

func conunt_num(i int) int {
	count := 0
	for i > 0{
		count++
		i = i & (i-1)
	}
	return count
}

func count_num2(n int) int {
	count := 0
	for n > 0{
		count += int(n & 1)
		n = n >> 1
	}
	return count
}
