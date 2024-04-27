// -*- coding:utf-8 -*-
// @Time : 2024/4/25 23:42
// @Author: MJ-CJM
// @File : leetcode/2739. 总行驶距离
package main

func distanceTraveled(mainTank int, additionalTank int) int {
	var res int = 0
	for mainTank > 0 {
		if mainTank >= 5 {
			res += 50
			mainTank -= 5
			if additionalTank > 0 {
				mainTank += 1
				additionalTank -= 1
			}
		} else {
			res += mainTank * 10
			mainTank = 0
		}
	}
	return res
}
