package main

import "sort"

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
