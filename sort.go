package main

import (
	"fmt"
	"math"
)

//冒泡排序，a是数组，n表示数组大小
func BubbleSort(a []int, n int) {
	if n <= 1 {
		return
	}
	for i := n - 1; i > 0; i-- {
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
			if a[j] < a[tmp] {
				tmp = j
			}
		}
		a[i], a[tmp] = a[tmp], a[i]
	}
	return a
}

// 快排
func QuickSort(a []int) []int {
	if len(a) <= 1 {
		return a
	} else {
		head := a[0]
		left, right := []int{}, []int{}
		for i := 1; i < len(a); i++ {
			if a[i] < head {
				left = append(left, a[i])
			} else {
				right = append(right, a[i])
			}
		}
		return append(append(QuickSort(left), head), QuickSort(right)...)
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
		} else {
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

// 堆排序
func HeapSort(a []int, n int) []int{
	buildheap(a, n)

	tmp := []int{}
	for i := 0; i < n; i++ {
		tmp = append(tmp, a[0])
		a = a[1:]
		buildheap(a, len(a))
	}
	tmp2 := reverse1(tmp)

	return tmp2
}

func reverse1(nums []int) []int {
	n := len(nums)
	if n <= 1{
		return nums
	}
	for i, j := 0, n-1; i < j; i,j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}

func buildheap(a []int, n int) {
	for i := n >> 1; i >= 0; i-- {
		heapifyUpToDown(a, i, n)
	}
	return
}

func heapifyUpToDown(a []int, top int, n int) {
	for i := top; i <= (n >> 1); {
		maxindex := i
		if (2 * i) < n && a[2 * i] > a[maxindex] {
			maxindex = 2 * i
		}
		if (2 * i + 1) < n && a[2 * i + 1] > a[maxindex] {
			maxindex = 2 * i + 1
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

// 计数排序
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

	// 构建C
	c := make([]int, max+1)
	for _, v := range a {
		c[v]++
	}
	for i := 1; i <= max; i++ {
		c[i] += c[i-1]
	}

	// 最终定位
	r := make([]int, n)
	for i := n-1; i >= 0; i-- {
		r[c[a[i]] - 1] = a[i]
		c[a[i]]--
	}
	copy(a, r)
}


func main() {
	a := []int{2, 4, 1, 3, 5}
	fmt.Println("before sort", a)
	// BubbleSort(a, len(a))
	// InsertionSort(a, len(a))
	// SelectSort(a, len(a))
	// res := QuickSort(a)
	// res := MergeSort(a)
	// res := HeapSort(a, len(a))
	CountingSort(a, len(a))
	fmt.Println("after sort", a)
}
