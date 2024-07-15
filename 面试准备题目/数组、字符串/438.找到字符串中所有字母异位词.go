package main

func findAnagrams(s string, p string) []int {
	n := len(s)
	m := len(p)
	var a[26]int
	match := 0
	for j := 0; j < m; j++ {
		a[p[j]-'a']++
	}
	var res []int
	for i := 0; i < n; i++ {
		a[s[i]-'a']--
		if a[s[i]-'a'] >= 0{
			match++
		}else{
			match--
		}
		if i >= m {
			a[s[i-m]-'a']++
			if a[s[i-m]-'a'] > 0 {
				match--
			}else{
				match++
			}
		}
		if match == m {
			res = append(res, i-m+1)
		}
	}
	return res
}

func findAnagrams2(s string, p string) []int {
	pE := getkey(p)
	n := len(s)
	m := len(p)
	res := []int{}
	if n < m {
		return res
	}
	for i := 0; i < n - m + 1; i++ {
		tmp := s[i:i+m]
		tmpK := getkey(tmp)
		if tmpK == pE {
			res = append(res, i)
		}
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