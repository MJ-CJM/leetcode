package main

import "fmt"

func main() {
	var c byte
	var err error
	var b []int
	for ; err == nil; {
		_, err = fmt.Scanf("%c", &c)

		if c != '\n' {
			if c != ' ' {
				b = append(b, int(c - '0'))
			}
		} else {
			break;
		}
	}
	res := majorityElement(b)
	fmt.Printf("%d", res)
}

func majorityElement(nums []int) int {
	major := 0
	count := 0
	for _, v := range nums{
		if count == 0{
			major = v
		}
		if major == v{
			count ++
		} else {
			count --
		}
	}
	return major
}