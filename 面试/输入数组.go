package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//var c byte
	//var err error
	//var b []int
	//for ; err == nil; {
	//	_, err = fmt.Scanf("%c", &c)
	//
	//	if c != '\n' {
	//		if c != ' ' {
	//			b = append(b, int(c - '0'))
	//		}
	//	} else {
	//		break;
	//	}
	//}
	//fmt.Println(b)

	var c byte
	var err error
	var b []int
	var d []string
	for ; err == nil; {
		_, err = fmt.Scanf("%c", &c)

		str := string(c)
		if str != "\n" {
			d = append(d, str)
		} else {
			break;
		}
	}
	s := strings.Join(d, "")
	s_l := strings.Split(s, " ")
	for _, v := range s_l {
		x, _ := strconv.Atoi(v)
		b = append(b, x)
	}
}