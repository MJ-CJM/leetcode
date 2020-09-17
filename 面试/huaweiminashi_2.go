package main

import "fmt"

func main() {
	hashmap := make(map[string]int)
	n := 10000
	var res string
	var s_l []string
	for i := 0; i < n; i++ {
		var s string
		_, _ = fmt.Scan(&s)
		if is_in_2(s_l, s) {
			continue
		}
		_, ok := hashmap[s]
		if ok {
			hashmap[s]++
			if s == res {
				res = ""
			}
		}else{
			hashmap[s] = 1
			if res != "" {
				res = s
			}
		}
		if hashmap[s] > 1 {
			s_l = append(s_l, s)
			delete(hashmap, s)
		}
	}
	fmt.Printf("%s", res)
}

func is_in_2(l []string, s string) bool {
	for _, v := range l {
		if v == s {
			return true
		}
	}
	return false
}