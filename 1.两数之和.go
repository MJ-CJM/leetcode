package main

func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int)
	for i, v := range nums {
		j, e := hashmap[target - v]
		if e == false {
			hashmap[v] = i
		}else{
			return []int{i,j}
		}
	}
	return nil
}
