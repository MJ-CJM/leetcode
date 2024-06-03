package main

func isPerfectSquare(num int) bool {
	if num <= 1 {
		return true
	}
	min, max := 0, num/2 + 1
	for min <= max {
		mid := (min + max) / 2
		if mid * mid > num {
			max = mid - 1
		}else if mid * mid < num {
			min = mid + 1
		}else{
			return true
		}
	}
	return false
}
