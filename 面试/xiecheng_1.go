package main

import (
	"fmt"
	"strings"
)

func main() {
	var m string
	var template string
	var t string
	_, _ = fmt.Scanln(&m)
	var c byte
	var err error
	var b []string
	count_1 := 0
	var jian []int
	for ; err == nil; {
		_, err = fmt.Scanf("%c", &c)

		if c != '\n' {
			if c != ' ' {
				b = append(b, string(c))
				count_1++
			}else{
				b = append(b, string(c))
				jian = append(jian, count_1)
				count_1++
			}
		} else {
			break;
		}
	}
	template = strings.Join(b, "")
	_, _ = fmt.Scanln(&t)
	result := process_tihuan(m, template, t)
	fmt.Printf("%s", result)
}

func process_tihuan(m string, template string, t string) string {
	m_l := []string{}
	tem_l := []string{}
	t_l := []string{}
	for _, v := range m {
		m_l = append(m_l, string(v))
	}
	for _, v := range template {
		tem_l = append(tem_l, string(v))
	}
	for _, v := range t {
		t_l = append(t_l, string(v))
	}
	n_m := len(m_l)
	//n_t := len(t_l)
	result := []string{}
	i := 0
	for i < len(tem_l) {
		temp := tem_l[i:i+n_m]
		if panduan_n_m(m, strings.Join(temp, "")) {
			result = append(result, t_l...)
			i += n_m
			continue
		}
		result = append(result, tem_l[i])
		i++
	}
	return strings.Join(result, "")
}

func panduan_n_m(m string, temp string) bool {
	s1 := make([]int, 27)
	s2 := make([]int, 27)
	for _, v := range m {
		if v >= 'a' && v <= 'z' {
			s1[v - 'a'] += 1
		}else{
			s1[26] += 1
		}
	}
	for _, v := range temp {
		if v >= 'a' && v <= 'z' {
			s2[v - 'a'] += 1
		}else{
			s2[26] += 1
		}
	}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
