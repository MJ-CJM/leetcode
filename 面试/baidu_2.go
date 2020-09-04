package main

import "fmt"

func main() {
	var zu int
	var n int
	var yueshu int
	var bing [][]int
	_, _ = fmt.Scan(&zu)
	_, _ = fmt.Scan(&n, &yueshu)
	for i := 0; i < yueshu; i++ {
		var count int
		var arr_tmp []int
		_, _ = fmt.Scan(&count)
		for j := 0; j < count; j++ {
			var left int
			var right int
			_, _ = fmt.Scan(&left, &right)
			for i := 1; i <= n; i++ {
				if i >= left && i <= right {
					arr_tmp = append(arr_tmp, i)
				}
			}
		}
		bing = append(bing, arr_tmp)
	}
	res, arr := hejie(bing, n)
	fmt.Printf("%d \n %v", res, arr)
}

func hejie(bing [][]int, n int) (int, []int) {
	var res []int
	for i := 1; i <= n; i++ {
		if is_in_bing(i, bing) {
			res = append(res, i)
		}
	}
	return len(res), res
}

func is_in_bing(x int, bing [][]int) bool {
	n := len(bing)
	for i := 0; i < n; i++ {
		if is_in_hang(x, bing[i]) {
			continue
		}
		return false
	}
	return true
}

func is_in_hang(x int, hang []int) bool {
	for _, v := range hang {
		if v == x {
			return true
		}
	}
	return false
}
