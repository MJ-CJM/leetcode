package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var c byte
	var err error
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
	res := process_pipei(d)
	n := len(res)
	for i, v := range res {
		if i == n-1 {
			fmt.Printf("%s", v)
			break
		}
		fmt.Printf("%s ", v)
	}
}

func process_pipei(d []string) []string {
	n := len(d)
	if n < 4 {
		return nil
	}
	res := []string{}
	for i := 0; i < n-3; i++ {
		if i == n-4 {
			if d[i] >= "0" && d[i] <= "9" {
				tmp := d[i:i+4]
				s := strings.Join(tmp, "")
				v, err := strconv.Atoi(s)
				if err == nil && v >= 1000 && v <= 3999 {
					res = append(res, s)
				}
			}
			continue
		}
		if i == 0 {
			if d[i] >= "0" && d[i] <= "9" {
				if d[i+4] >= "0" && d[i+4] <= "9" {
					continue
				}
				tmp := d[i:i+4]
				s := strings.Join(tmp, "")
				v, err := strconv.Atoi(s)
				if err == nil && v >= 1000 && v <= 3999 {
					res = append(res, s)
				}
			}
		}
		if d[i] >= "0" && d[i] <= "9" {
			if d[i+4] >= "0" && d[i+4] <= "9" {
				continue
			}
			if d[i-1] >= "0" && d[i-1] <= "9" {
				continue
			}
			tmp := d[i:i+4]
			s := strings.Join(tmp, "")
			v, err := strconv.Atoi(s)
			if err == nil && v >= 1000 && v <= 3999 {
				res = append(res, s)
			}
		}
	}
	return res
}

//import re
//
//def findit(string):
//list1 = re.findall(r"\d+",string)
//for i in range(len(list1)):
//if i == len(list1) - 1 and 1000 <= int(list1[i]) <= 3999:
//print(list[i])
//continue
//if 1000 <= int(list1[i]) <= 3999:
//print(list1[i], end=" ")
//
//string = input()
//findit(string)
