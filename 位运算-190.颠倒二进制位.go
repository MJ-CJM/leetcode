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
