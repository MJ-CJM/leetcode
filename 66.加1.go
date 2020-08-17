package main

func plusOne(digits []int) []int {
	n := len(digits)
	if n == 0{
		return []int{1}
	}
	for i := n-1; i >= 0; i-- {
		if digits[i] != 9{
			digits[i]++
			return digits
		}else{
			digits[i] = 0
		}
	}
	tmp := append([]int{1}, digits...)
	return tmp
}
