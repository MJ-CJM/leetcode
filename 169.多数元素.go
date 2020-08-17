package main

// 暴力解法
func majorityElement1(nums []int) int {
	var result int
	tmp := len(nums) / 2
	hashmap := make(map[int]int)
	for _, v := range nums {
		if _, e := hashmap[v]; e != true{
			hashmap[v] = 1
		}else{
			hashmap[v] += 1
		}
	}
	for k, v := range hashmap{
		if v > tmp {
			result = k
		}
	}
	return result
}

// 摩尔投票法
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
