package main

import (
	"fmt"
)

func countSmaller(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	if n == 0 {
		return res
	}

	// 扩展数组，存储值和索引
	indexedNums := make([][2]int, n)
	for i, num := range nums {
		indexedNums[i] = [2]int{num, i}
	}

	// 归并排序并统计
	mergeSort(indexedNums, res)

	return res
}

func mergeSort(nums [][2]int, res []int) [][2]int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	left := mergeSort(nums[:mid], res)
	right := mergeSort(nums[mid:], res)

	return merge(left, right, res)
}

func merge(left, right [][2]int, res []int) [][2]int {
	merged := make([][2]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i][0] <= right[j][0] {
			merged = append(merged, left[i])
			res[left[i][1]] += j // 左边元素右边比它小的个数就是 j
			i++
		} else {
			merged = append(merged, right[j])
			j++
		}
	}

	for i < len(left) {
		merged = append(merged, left[i])
		res[left[i][1]] += j // 左边剩余元素右边比它小的个数也是 j
		i++
	}

	for j < len(right) {
		merged = append(merged, right[j])
		j++
	}

	return merged
}

func main() {
	nums := []int{5, 2, 6, 1}
	res := countSmaller(nums)
	fmt.Println(res) // 输出: [2, 1, 1, 0]
}
