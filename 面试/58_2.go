package main

/**
 *
 * @param num int整型 非负整数 num
 * @return int整型一维数组
 */
func countBits( num int ) []int {
	// write code here
	res := []int{}
	for i := 0; i <= num; i++ {
		tmp := hammingWeight(uint32(i))
		res = append(res, tmp)
	}
	return res
}

func hammingWeight(num uint32) int {
	count := 0
	for num > 0{
		count++
		num = num & (num-1)
	}
	return count
}
