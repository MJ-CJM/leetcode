package main

import "strings"

func permutation(s string) []string {
	l := strings.Split(s, "")
	res := []string{}
	dfs_per(l, len(l), []string{}, &res)
	return res
}

func dfs_per(l []string, n int, iterm []string, res *[]string) {
	// terminator
	if len(iterm) == n {
		tmp := strings.Join(iterm, "")
		if !isin_string(tmp, *res) {
			*res = append(*res, tmp)
		}
		return
	}
	// process && drill down
	for i, v := range l {
		c := append(iterm, v)
		tmp2 := make([]string, len(l)-1)
		copy(tmp2[:i], l[:i])
		copy(tmp2[i:], l[i+1:])
		dfs_per(tmp2, n, c, res)
	}
}

func isin_string(tmp string, res []string) bool {
	for _, v := range res {
		for i := 0; i < len(tmp); i++ {
			if v[i] != tmp[i] {
				break
			}
			if i == len(tmp)-1 {
				return true
			}
		}
	}
	return false
}


