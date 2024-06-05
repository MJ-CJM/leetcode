package main

func partitionLabels(s string) []int {
	// 记录每个字符最后出现的位置
	last := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		last[s[i]] = i
	}

	var result []int
	start, end := 0, 0

	for i := 0; i < len(s); i++ {
		// 更新当前片段的结束位置
		if last[s[i]] > end {
			end = last[s[i]]
		}
		// 当到达片段的结束位置时，记录片段长度
		if i == end {
			result = append(result, end-start+1)
			start = i + 1
		}
	}

	return result
}
