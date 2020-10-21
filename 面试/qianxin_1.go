package main

import "fmt"

/**
 * 代码中的类名、方法名、参数名已经指定，请勿修改，直接返回方法规定的值即可
 * M包糖果，抛M次硬币，硬币连续n次为正面，最多能得到多少颗糖果
 * @param candies int整型一维数组 每包糖果的数量
 * @param coin int整型一维数组 抛硬币的结果
 * @param n int整型 连续是正面的次数
 * @return int整型
 */
func maxCandies( candies []int ,  coin []int ,  n int ) int {
	// write code here
	m := len(candies)
	if n >= m {
		res := 0
		for _, v := range candies {
			res += v
		}
		return res
	}
	target := 0
	for i := 0; i <= m - n; i++ {
		tmp2 := 0
		tmp := make([]int, len(coin))
		copy(tmp, coin)
		for j := i; j < i + n; j++ {
			tmp[j] = 0
		}
		for k := 0; k < len(tmp); k++ {
			if tmp[k] == 0 {
				tmp2 += candies[k]
			}
		}
		if tmp2 > target {
			target = tmp2
		}
	}
	return target
}

func main() {
	candi := []int{3,5,7,2,8,8,15,3}
	coin := []int{1,0,1,0,1,0,1,0}
	count := 3
	fmt.Println(maxCandies(candi, coin, count))
}