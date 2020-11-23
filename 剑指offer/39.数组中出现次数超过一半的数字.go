package main

func majorityElement(nums []int) int {
	maps := make(map[int]int)
	n := len(nums)
	for _, v := range nums {
		_, ok := maps[v]
		if ok {
			maps[v]++
			if maps[v] > (n>>1) {
				return v
			}
		}else{
			maps[v] = 1
		}
	}
	return nums[0]
}
