package main

// 移动32次
func hammingWeight(num uint32) int {
	count := 0
	for num > 0 {
		count += int(num & 1)
		num = num >> 1
	}
	return count
}

// 清零最低位的1
func hammingWeight2(num uint32) int {
	count := 0
	for num > 0{
		count++
		num = num & (num-1)
	}
	return count
}