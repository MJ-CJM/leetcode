package main

func numJewelsInStones(J string, S string) int {
	if S == ""{
		return 0
	}
	count := 0
	for _, v := range S {
		for _, k := range J {
			if v == k {
				count++
				break
			}
		}
	}
	return count
}