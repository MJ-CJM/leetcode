package main

func lemonadeChange(bills []int) bool {
	hash := make(map[int]int)
	hash[5] = 0
	hash[10] = 0
	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			hash[5]++
			continue
		}
		if bills[i] == 10 {
			if hash[5] > 0 {
				hash[5]--
				hash[10]++
			}else{
				return false
			}
			continue
		}
		if bills[i] == 20 {
			if hash[5] > 0 && hash[10] > 0 {
				hash[5]--
				hash[10]--
			}else if hash[5] > 2{
				hash[5] -= 3
			}else{
				return false
			}
			continue
		}
	}
	return true
}
