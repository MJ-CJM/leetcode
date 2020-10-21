package main

func minArray(numbers []int) int {
	n := len(numbers)
	left := 0
	right := n-1
	for left < right {
		mid := left + (right - left) >> 1
		if numbers[mid]  > numbers[right] {
			left = mid + 1
		}else{
			right = right - 1
		}
	}
	return numbers[left]
}
