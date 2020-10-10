package main

func reverseBits(num uint32) uint32 {
	var res uint32 = 0
	for i := uint32(0); i < 32; i++ {
		tmp := num & 1
		res += tmp << (32 - i -1)
		num = num >> 1
	}
	return res
}

/*将给定的uint32数字进行二进制反转， 只需按2取余（r）即可
从最高位开始，根据余数r，进行位运算。
将每次位运算得到的值和结果值进行或运算*/

func reverseBits_1(num uint32) uint32 {
	var ans,zero uint32
	var bits uint32 = 31

	for num > 0 {
		r := num%2
		ans = ans | (zero+r)<<bits
		bits--
		num = num/2
	}
	return ans
}
