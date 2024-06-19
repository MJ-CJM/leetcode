package main

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	m := make(map[[26]int][]string)
	for _, k := range strs {
		key := getkey(k)
		_, ok := m[key]
		if ok {
			m[key] = append(m[key], k)
		}else{
			m[key] = []string{k}
		}
	}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

func getkey(s string) [26]int {
	res := [26]int{}
	for i := 0; i < len(s); i++ {
		res[s[i]-'a']++
	}
	return res
}