package main

import (
	"fmt"
	"math"
)

func count(nums []bool) int {
	res := 0
	for _, v := range nums {
		if v == true {
			res++
		}
	}
	return res
}

/*
3
5 6 10
1
1 2

*/

func main() {
	var appNum int
	fmt.Scan(&appNum)

	runTimes := make([]int, appNum)
	for i := 0; i < appNum; i++ {
		fmt.Scan(&runTimes[i])
	}

	var mutexNum int
	fmt.Scan(&mutexNum)

	// 创建一个二维数组来表示互斥关系
	mutex := make([][]bool, appNum)
	for i := range mutex {
		mutex[i] = make([]bool, appNum)
	}

	for i := 0; i < mutexNum; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		a-- // 转为0索引
		b--
		mutex[a][b] = true
		mutex[b][a] = true
	}

	minTime := math.MaxInt32
	maxCount := 0

	// 回溯算法，递归搜索最小运行时间
	var backtrack func(index int, currentTime int, selected []bool)
	backtrack = func(index int, currentTime int, selected []bool) {
		if index == appNum {
			if count(selected) == 0 {
				return
			}
			if count(selected) > maxCount {
				minTime = currentTime
				maxCount = count(selected)
			} else if count(selected) == maxCount {
				if currentTime < minTime {
					minTime = currentTime
				}
			}
			return
		}

		// 选择不包含当前程序
		backtrack(index+1, currentTime, selected)

		// 检查是否可以包含当前程序
		canSelect := true
		for i := 0; i < index; i++ {
			if selected[i] && mutex[i][index] {
				canSelect = false
				break
			}
		}

		if canSelect {
			// 选择包含当前程序
			selected[index] = true
			backtrack(index+1, currentTime+runTimes[index], selected)
			selected[index] = false
		}
	}

	// 初始化回溯
	selected := make([]bool, appNum)
	backtrack(0, 0, selected)

	// 如果 minTime 仍然是初始值，说明没有找到有效路径
	if minTime == math.MaxInt32 {
		minTime = 0
	}

	fmt.Println("res: ", minTime)
}

