package main

import "fmt"

// 动态规划：重复子问题:第i天的最大值
// 定义状态：mp[i][k][j], j = 0/1 0:没有股票，1：有一股股票,k交易了多少次
// 状态转移方程：mp[i][k][0] = max(mp[i-1][k][0], mp[i-1][k-1][1]+a[i]//卖出)
//             mp[i][k][1] = max(mp]i-1][k][1], mp[i-1][k-1][0]-a[i]//买入)
// 最后求：max(mp[n-1][k][0])
func maxProfit4(k int, prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}
	if k >= (n / 2){
		return maxProfit2(prices)
	}
	mp := make([][][]int, n)
	for i := 0; i < n; i++ {
		mp[i] = make([][]int, k+1)
		for j := 0; j < k+1; j++ {
			mp[i][j] = make([]int, 2)
		}
	}
	for i := 0; i < n; i++ {
		for j := k; j >= 1; j-- {
			if i == 0 {
				mp[i][j][0] = 0
				mp[i][j][1] = -prices[i]
				continue
			}
			mp[i][j][0] = max(mp[i-1][j][0], mp[i-1][j][1]+prices[i])
			mp[i][j][1] = max(mp[i-1][j][1], mp[i-1][j-1][0]-prices[i])
		}
	}
	return mp[n-1][k][0]
}

func max_1(x, y int) int {
	if x > y {
		return x
	}else{
		return y
	}
}

func main(){
	k := 2
	price := []int{2, 4, 1}
	fmt.Println(maxProfit4(k, price))
}