package main

import "fmt"

func main() {
	var n int
	var p int
	_, _ = fmt.Scan(&n, &p)
	l_count := []int{}
	l_price := []int{}
	l_meili := []int{}
	for i := 0; i < n; i++ {
		x, y, z := 0, 0, 0
		_, _ = fmt.Scan(&x, &y, &z)
		l_count = append(l_count, x)
		l_price = append(l_price, y)
		l_meili = append(l_meili, z)
	}
	res := process_maxmeili(n, p, l_count, l_price, l_meili)
	fmt.Printf("%d", res)
}

func process_maxmeili(n int, p int, count []int, price []int, meili []int) int {
	hashmap := make(map[int]int)
	for i := 0; i < n; i++ {
		hashmap[price[i]] = meili[i]
	}
	jiage := []int{}
	for i := 0; i < n; i++ {
		for j := 0; j < count[i]; j++ {
			jiage = append(jiage, price[i])
		}
	}
	iterm := []int{}
	level := 0
	res := 0
	_dfs_meili(p, level, iterm, jiage, hashmap, &res)
	return res
}

func _dfs_meili(p int, level int, iterm []int, jiage []int, hashmap map[int]int,res *int) {
	// terminator
	if count_int(iterm) > p {
		return
	}
	if count_int(iterm) == p {
		tmp := 0
		for _, v := range iterm {
			tmp += hashmap[v]
		}
		if tmp > *res {
			*res = tmp
		}
		return
	}
	if level >= len(jiage) {
		he := count_int(iterm)
		if he <= p {
			tmp := 0
			for _, v := range iterm {
				tmp += hashmap[v]
			}
			if tmp > *res {
				*res = tmp
			}
		}
		return
	}
	// process && drill down
	c1 := iterm
	c2 := append(iterm, jiage[level])
	_dfs_meili(p, level+1, c1, jiage, hashmap, res)
	_dfs_meili(p, level+1, c2, jiage, hashmap, res)
}

func count_int(nums []int) int {
	n := len(nums)
	res := 0
	for i := 0; i < n; i++ {
		res += nums[i]
	}
	return res
}
