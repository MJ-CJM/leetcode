package 数组_字符串

func numJewelsInStones(jewels string, stones string) int {
	jMap := make(map[rune]bool)
	for _, v := range jewels {
		jMap[v] = true
	}
	res := 0
	for _, v := range stones {
		if jMap[v] == true {
			res++
		}
	}
	return res
}