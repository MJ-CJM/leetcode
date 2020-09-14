package main

import "fmt"

/**
 * 找出字符串中最大公共子字符串
 * @param str1 string字符串 输入字符串1
 * @param str2 string字符串 输入字符串2
 * @return string字符串
 */
func GetCommon( str1 string ,  str2 string ) string {
	n := len(str1)
	m := len(str2)
	res := ""
	for i := 0; i < n; i++ {
		pre_i := i
		tmp := ""
		for j := 0; j < m; j++ {
			if str1[i] == str2[j] && i < n && j < m {
				tmp += string(str1[i])
				i++
				continue
			}
			if len(tmp) > 0 {
				break
			}
		}
		if len(tmp) > len(res) {
			res = tmp
		}
		i = pre_i
	}
	return res
}

func main() {
	s := "abccade"
	t := "dgcadde"
	fmt.Println(GetCommon(s, t))
}
