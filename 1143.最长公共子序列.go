package main

import "fmt"

// 暴力算法求解
//func longestCommonSubsequence(text1 string, text2 string) int {
//	s1 := subsequence(text1)
//	s2 := subsequence(text2)
//	result := 0
//	for _, v1 := range s1{
//		for _, v2 := range s2 {
//			if v1 == v2{
//				if len(v1) > result{
//					result = len(v1)
//				}
//			}
//		}
//	}
//	return result
//}
//
//func subsequence(s string) []string{
//	result := []string{}
//	n := len(s)
//	level := 0
//	iterm := ""
//	_sub(s, n, level, iterm, &result)
//	return result
//}
//
//func _sub(s string, n int, i int, iterm string, result *[]string) {
//	// terminator
//	if i == n{
//		*result = append(*result, iterm)
//		return
//	}
//
//	// process && drill in
//	c1 := iterm
//	c2 := iterm
//	_sub(s, n, i+1, c1, result)
//	c2 += string(s[i])
//	_sub(s, n, i+1, c2, result)
//}

func longestCommonSubsequence(text1 string, text2 string) int {
	n := len(text1)
	m := len(text2)
	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, n)
	}
	temp := 0
	for i := 0; i < n; i++ {
		if i == 0 && text1[i] == text2[0] {
			temp = -1
			break
		}
		if text1[i] == text2[0] {
			temp = i
			break
		}
	}
	if temp != 0 {
		if temp == -1 {
			temp = 0
			for i := temp; i < n; i++ {
				result[0][i] = 1
			}
		} else {
			for i := temp; i < n; i++ {
				result[0][i] = 1
			}
		}
	}
	temp2 := 0
	for i := 0; i < m; i++ {
		if i == 0 && text2[i] == text1[0] {
			temp2 = -1
			break
		}
		if text2[i] == text1[0] {
			temp2 = i
			break
		}
	}
	if temp2 != 0 {
		if temp2 == -1 {
			temp2 = 0
			for i := temp2; i < m; i++ {
				result[i][0] = 1
			}
		} else {
			for i := temp2; i < m; i++ {
				result[i][0] = 1
			}
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if text2[i] == text1[j]{
				result[i][j] = result[i-1][j-1] + 1
			}else{
				result[i][j] = max_int(result[i-1][j], result[i][j-1])
			}
		}
	}
	return result[m-1][n-1]
}

func max_int(x int, y int)int{
	if x >= y{
		return x
	}else{
		return y
	}
}

func main() {
	s1 := "bsbininm"
	s2 := "jmjkbkjkv"
	fmt.Println(longestCommonSubsequence(s1, s2))
}
