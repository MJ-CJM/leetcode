package main

func maxProfit2(prices []int) int {
	n := len(prices)
	result := 0
	for i := 0; i < n-1; i++ {
		if prices[i+1] > prices[i] {
			result += prices[i+1] - prices[i]
		}
	}
	return result
}


