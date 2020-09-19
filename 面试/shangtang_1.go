package main

func main() {
	//var err error
	//n, m := 1, 1
	//res := []string{}
	//for ; err == nil; {
	//	_, err = fmt.Scan(&n, &m)
	//	if n == 0 && m == 0{
	//		break
	//	}
	//	tmp := []string{}
	//	for i := 0; i < n; i++ {
	//		var s string
	//		_, _ = fmt.Scan(&s)
	//		tmp = append(tmp, s)
	//	}
	//	output := process_migong(tmp, n, m)
	//	res = append(res, output)
	//}
	//for _, v := range res {
	//	fmt.Printf("%s\n", v)
	//}
}

func process_migong(tmp []string, n int, m int) string {
	x, y := 0, 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if tmp[i][j] == 'S' {
				x, y = i, j
				break
			}
		}
	}
	hashmap := make(map[byte]int)
	res := "NO"
	dfs_migong(tmp, x, y, hashmap, &res)
	return res
}

func dfs_migong(tmp []string, x int, y int, hashmap map[byte]int, res *string)  {
	// terminator
	if x < 0 || x >= len(tmp) || y < 0 || y >= len(tmp[0]) {
		return
	}
	if tmp[x][y] == 'G' {
		*res = "YES"
		return
	}
	if tmp[x][y] == 'X' {
		return
	}

	// process && drill down
	if tmp[x][y] ==  'a' {
		hashmap['a'] += 1
	}
	if tmp[x][y] ==  'b' {
		hashmap['b'] += 1
	}
	if tmp[x][y] ==  'c' {
		hashmap['c'] += 1
	}
	if tmp[x][y] ==  'd' {
		hashmap['d'] += 1
	}
	if tmp[x][y] ==  'e' {
		hashmap['e'] += 1
	}
	if tmp[x][y] ==  'A' {
		v, ok := hashmap['a']
		if ok {
			if v == 0 {
				return
			}
			hashmap['a'] -= 1
		}else{
			return
		}
	}
	if tmp[x][y] ==  'B' {
		v, ok := hashmap['b']
		if ok {
			if v == 0 {
				return
			}
			hashmap['b'] -= 1
		}else{
			return
		}
	}
	if tmp[x][y] ==  'C' {
		v, ok := hashmap['c']
		if ok {
			if v == 0 {
				return
			}
			hashmap['c'] -= 1
		}else{
			return
		}
	}
	if tmp[x][y] ==  'D' {
		v, ok := hashmap['d']
		if ok {
			if v == 0 {
				return
			}
			hashmap['d'] -= 1
		}else{
			return
		}
	}
	if tmp[x][y] ==  'E' {
		v, ok := hashmap['e']
		if ok {
			if v == 0 {
				return
			}
			hashmap['e'] -= 1
		}else{
			return
		}
	}
	dfs_migong(tmp, x+1, y, hashmap, res)
	dfs_migong(tmp, x-1, y, hashmap, res)
	dfs_migong(tmp, x, y+1, hashmap, res)
	dfs_migong(tmp, x, y-1, hashmap, res)
}
