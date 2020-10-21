package main

func validateStackSequences(pushed []int, popped []int) bool {
	zhan := []int{}
	i := 0
	for _, v := range pushed{
		zhan = append(zhan, v)
		for len(zhan) != 0 && zhan[len(zhan)-1] == popped[i] {
			zhan = zhan[:len(zhan)-1]
			i++
		}
	}
	if len(zhan) == 0 {
		return true
	}else{
		return false
	}
}
