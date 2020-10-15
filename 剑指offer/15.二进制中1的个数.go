package main

func hammingWeight(num uint32) int {
	count := 0
	for num > 0 {
		count += int(num & 1)
		num = num >> 1
	}
	// 清掉最低位的0
	//for num > 0 {
	//	count++
	//	num = num & (num-1)
	//}
	return count
}
