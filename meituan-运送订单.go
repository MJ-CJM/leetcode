package main

import (
	"fmt"
	"sort"
)

func main(){
	var n, m int
	_, _ = fmt.Scan(&n, &m)
	hm := make(map[int]int, 0)
	for i := 0; i < m; i++ {
		var x, y int
		_, _ = fmt.Scan(&x, &y)
		hm[x] = y
	}
	list := peisong(n, hm)
	count := len(list)
	fmt.Println(count)
	for i := 0; i < count; i++ {
		//list[i] = quchong(list[i])
		for j := 0; j < len(list[i]); j++ {
			if j == len(list[i]) - 1{
				fmt.Printf("%d\n", list[i][j])
				break
			}
			fmt.Printf("%d ", list[i][j])
		}
	}
}

func quchong(list []int) []int {
	res := []int{}
	for i := 0; i < len(list); i++ {
		if (i > 0 && list[i-1] == list[i]) {
			continue
		}
		res = append(res, list[i])
	}
	return res
}

func peisong(n int, hm map[int]int) [][]int {
	res := [][]int{}
	dif := []int{}
	for i := 1; i <= n; i++ {
		v, ok := hm[i]
		if ok {
			res = v_append(res, i, v)
			continue
		}
		dif = append(dif, i)
	}
	for i := 0; i < len(dif); i++ {
		res = dif_append(res, dif[i])
	}
	for i := 0; i < len(res); i++ {
		sort.Ints(res[i])
	}
	return res
}

func dif_append(res [][]int, x int) [][]int {
	n := len(res)
	flag := false
	for i := 0; i < n; i++ {
		flag, res[i] = is_in_list(res[i], x, x)
		if flag {
			break
		}
	}
	if !flag {
		res = append(res, []int{x})
	}
	return res
}

func v_append(res [][]int, i int, v int) [][]int {
	n := len(res)
	flag := false
	for j := 0; j < n; j++ {
		flag, res[j] = is_in_list(res[j], i, v)
		if flag {
			break
		}
	}
	if !flag {
		res = append(res, []int{i, v})
	}
	return res
}

func is_in_list(list []int, x int, y int) (bool, []int) {
	n := len(list)
	flag := false
	for i := 0; i < n; i++ {
		if list[i] == x || list[i] == y {
			list = append(list, []int{x, y}...)
			flag = true
			break
		}
	}
	return flag, list
}
