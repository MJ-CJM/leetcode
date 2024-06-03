package main

import (
	"fmt"
)

//func minMutation(start string, end string, bank []string) int {
//	used := make([]int, len(bank))
//	queue := []string{start}
//	num := 0
//	for len(queue) != 0{
//		num ++
//		level := len(queue)
//		for i := 0; i < level; i ++{
//			for k, v := range bank{
//				if v != start && used[k] == 0{
//					diff := 0
//					for j := 0; j < 8; j ++{
//						if v[j] != queue[i][j]{
//							diff ++
//						}
//						if diff >= 2{
//							break
//						}
//					}
//					if diff == 1{
//						used[k] = num
//						if bank[k] == end{
//							return num
//						}
//						queue = append(queue, bank[k])
//					}
//				}
//			}
//		}
//		queue = queue[level:]
//	}
//	return -1
//}

var wMap = map[byte][]string{
	'A': []string{"C", "G", "T"},
	'C': []string{"A", "G", "T"},
	'G': []string{"A", "C", "T"},
	'T': []string{"A", "C", "G"},
}

var (
	used     []bool
	minCount int
)

func minMutation(start string, end string, bank []string) int {
	if i := isIn(bank, end); i == -1 {
		return -1
	}
	minCount = len(bank) + 1
	used = make([]bool, len(bank))
	_dfs(start, end, bank, 0)
	if minCount <= len(bank) {
		return minCount
	}
	return -1
}

func _dfs(temp string, end string, bank []string, counter int) {
	// terminator
	if counter >= minCount {
		return
	}
	if temp == end {
		if counter < minCount {
			minCount = counter
		}
		return
	}
	// process && drill down
	for i := range temp {
		for _, n := range wMap[temp[i]] {
			next := temp[:i] + n + temp[i+1:]
			idx := isIn(bank, next)
			if idx != -1 && !used[idx] {
				used[idx] = true
				_dfs(next, end, bank, counter+1)
				// reverse states
				used[idx] = false
			}
		}
	}
}

func isIn(bank []string, x string) int {
	for i, v := range bank {
		if x == v {
			return i
		}
	}
	return -1
}

func main() {
	start := "AACCGGTT"
	end := "AACCGGTA"
	bank := []string{"AACCGGTA"}
	fmt.Println(minMutation(start, end, bank))
}
