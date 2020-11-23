package main

func firstUniqChar(s string) byte {
	tmp := [26]int{}
	for _, v := range s {
		tmp[v-'a']++
	}
	for _, v := range s {
		if tmp[v-'a'] == 1 {
			return byte(v)
		}
	}
	return ' '
}
