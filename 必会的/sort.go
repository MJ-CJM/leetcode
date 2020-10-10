package main

import "math"

//冒泡排序, n*n, 稳定， 原地
func BubbleSort(a []int, n int) {
	if n <= 1 {
		return
	}
	for i := n-1; i > 0; i-- {
		flag := 0
		for j := 0; j < i; j++ {
			if a[j] > a[j+1] {
				flag = 1
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
		if flag == 0 {
			break
		}
	}
	return
}

// 插入排序, n*n, 稳点， 原地
func InsertionSort(a []int, n int) {
	if n <= 1 {
		return
	}
	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
	return
}

// 选择排序, n*n, 不稳定, 原地
func SelectSort(a []int, n int) []int {
	if n <= 1 {
		return a
	}
	for i := 0; i < n; i++ {
		tmp := i
		for j := i; j < n; j++ {
			if a[tmp] > a[j] {
				tmp = j
			}
		}
		a[tmp], a[i] = a[i], a[tmp]
	}
	return a
}

// 快排, nlogn, 不稳定, 原地
func QuickSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}else{
		mid := a[0]
		left := []int{}
		right := []int{}
		for i := 1 ;i < len(a); i++ {
			if a[i] < mid {
				left = append(left, a[i])
			}else{
				right = append(right, a[i])
			}
		}
		return append(append(QuickSort(left), mid), QuickSort(right)...)
	}
}

// 两路归并排序, nlogn, 稳定, 非原地
func MergeSort(a []int) []int {
	if len(a) <= 1 {
		return a
	}
	mid := len(a) >> 1
	left := MergeSort(a[:mid])
	right := MergeSort(a[mid:])
	return Merge(left, right)
}

func Merge(left []int, right []int) []int {
	var res []int
	l := len(left)
	r := len(right)
	i := 0
	j := 0
	for i < l && j < r {
		if left[i] < right[j] {
			res = append(res, left[i])
			i++
		}else{
			res = append(res, right[j])
			j++
		}
	}
	if i < l {
		res = append(res, left[i:]...)
	}
	if j < r {
		res = append(res, right[j:]...)
	}
	return res
}

// 堆排序, nlogn, 不稳定， 原地
func HeapSort(a []int, n int) []int {
	buildheap(a, n)

	for i := 0; i < n; i++ {
		swap(a, 0, n-i-1)
		buildheap(a[:n-i-1], n-i-1)
	}

	return a
}

func buildheap(a []int, n int) {
	for i := n >> 1; i >= 0; i-- {
		heapifyUpToDown(a, i, n)
	}
	return
}

func heapifyUpToDown(a []int, top int, n int) {
	for i := top; i <= n >> 1; {
		maxindex := i
		if 2 * i < n && a[2 * i] > a[maxindex] {
			maxindex = 2 * i
		}
		if 2 * i + 1 < n && a[2 * i +1] > a[maxindex] {
			maxindex = 2 * i +1
		}
		if maxindex == i {
			break
		}
		swap(a, i, maxindex)
		i = maxindex
	}
}

func swap(a []int, i int, j int) {
	a[i], a[j] = a[j], a[i]
}

// 计数排序, n+k, 稳定, 非原地
func CountingSort(a []int, n int) {
	if n <= 1 {
		return
	}

	// 找出最大值
	var max = math.MinInt32
	for i := 0; i < n; i++ {
		if a[i] > max {
			max = a[i]
		}
	}

	// 构建c
	c := make([]int, max+1)
	for _, v := range a {
		c[v]++
	}
	for i := 1; i <= max; i++ {
		c[i] += c[i-1]
	}

	// 最终定位
	r := make([]int, n)
	for i := n-1; i>=0; i-- {
		r[c[a[i]]-1] = a[i]
		c[a[i]]--
	}
	copy(a, r)
}
