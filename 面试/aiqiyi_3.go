package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
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
	res := threeSum(b)
	n := len(res)
	m := len(res[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%d ", res[i][j])
		}
		fmt.Println()
	}
}

func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	output := make([][]int, 0)

	for i := 0; i < n-1; i++ {
		p := 0
		q := n - 1
		if i > 0 && nums[i] == nums[i-1]{
			p = i-1
		}
		for p < i && i < q{
			tmp := nums[p] + nums[i] + nums[q]
			if p > 0 && nums[p] == nums[p-1]{
				p ++
				continue
			}
			if q < n-1 && nums[q] == nums[q+1]{
				q --
				continue
			}
			if tmp > 0{
				q --
				continue
			}else if tmp < 0{
				p ++
				continue
			}else{
				output = append(output, []int{nums[p], nums[i], nums[q]})
				p ++
				q --
				continue
			}
		}
	}
	return output
}

