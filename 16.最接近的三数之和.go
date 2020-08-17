package main

import (
	"fmt"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	sort.Ints(nums)
	cha := nums[0] + nums[1] + nums[2] - target
	if cha < 0 {
		cha = -cha
	}
	output := nums[0] + nums[1] + nums[2]

	for i := 0; i < n - 1; i++ {
		p := 0
		q := n - 1
		if i > 0 && nums[i] == nums[i-1]{
			p = i-1
		}
		for p < i && i < q {
			tmp := nums[p] + nums[i] + nums[q]
			if p > 0 && nums[p] == nums[p-1]{
				p ++
				continue
			}
			if q < n - 1 && nums[q] == nums[q+1]{
				q --
				continue
			}
			if tmp < target{
				process(tmp, target, &cha, &output)
				p ++
				continue
			}else if tmp > target{
				process(tmp, target, &cha, &output)
				q --
				continue
			}else{
				process(tmp, target, &cha, &output)
				p ++
				q --
				continue
			}
		}
	}
	return output
}

func process(x int, target int, cha *int, output *int) {
	tmp := x - target
	if tmp < 0{
		tmp = -tmp
	}
	if tmp < *cha{
		*cha = tmp
		*output = x
	}
}

func main () {
	s := []int{-1, 2, 1, -4}
	target := 1
	output := threeSumClosest(s, target)
	fmt.Println(output)
}