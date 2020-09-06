package main

//冒泡排序，a是数组，n表示数组大小
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

// 插入排序，a表示数组，n表示数组大小
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

// 选择排序
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

// 快排
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

// 两路归并排序
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