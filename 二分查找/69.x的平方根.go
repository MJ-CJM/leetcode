package main

func mySqrt(x int) int {
	if x <= 1{
		return x
	}
	left := 0
	right := x
	mid := 0
	for left <= right{
		mid = left + (right - left)/2
		if mid*mid <= x && (mid+1)*(mid+1) > x{
			return mid
		}else if mid * mid > x{
			right = mid -1
		}else {
			left = mid + 1
		}
	}
	return mid
}

// 牛顿迭代法
func NewDun(x int) int {
	r := x
	for r * r > x{
		r = (r + x/r) / 2
	}
	return r
}
