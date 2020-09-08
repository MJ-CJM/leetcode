package main

import "fmt"

func main() {
	n := 0
	m := 0
	_, _ = fmt.Scan(&n, &m)
	tmp := [][]int{}
	for i := 0; i < m; i++ {
		m_len := 0
		_, _ = fmt.Scan(&m_len)
		zi := []int{}
		for i := 0; i < m_len; i++ {
			x := 0
			_, _ = fmt.Scan(&x)
			zi = append(zi, x)
		}
		tmp = append(tmp, zi)
	}
	res := []int{}
	process_tongzhi(tmp, &res)
	fmt.Printf("%d", len(res))
}

func process_tongzhi(tmp [][]int, res *[]int) int {
	//output := []int{}
	//for i := 0; i < len(tmp); i++ {
	//	if is_in(tmp[i], 0) {
	//		for j := 0; j < len(tmp[i]); j++ {
	//			output = append(output, tmp[i][j])
	//		}
	//		tmp = append(tmp[:i], tmp[i+1:]...)
	//	}
	//}
	// terminator
	return 0
}

func is_in(nums []int, x int) bool {
	for _, v := range nums {
		if v == x {
			return true
		}
	}
	return false
}


