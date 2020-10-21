package main

import "fmt"

func main() {
	m, n := 0, 0
	_, _ = fmt.Scan(&m, &n)
	res := process_juzheng(m, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%s ", res[i][j])
		}
		fmt.Println()
	}
}

func process_juzheng(m int, n int) [][]string {
	res := make([][]string, m)
	for i := 0; i < m; i++ {
		res[i] = make([]string, n)
	}
	tmp := make([][]byte, m)
	for i := 0; i < m; i++ {
		tmp[i] = make([]byte, n)
	}
	res[0][0] = "A"
	tmp[0][0] = 'A'
	top := 0
	bottom := m - 1
	left := 0
	right := n - 1
	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			if i == 0 && top == 0 {
				tmp[0][0] = 'A'
				continue
			}
			if tmp[top][i-1] == 'Z' {
				tmp[top][i] = 'A'
				continue
			}
			tmp[top][i] = tmp[top][i-1] + 1
		}
		for i := top+1; i <= bottom; i++ {
			if tmp[i-1][right] == 'Z' {
				tmp[i][right] = 'A'
				continue
			}
			tmp[i][right] = tmp[i-1][right] + 1
		}
		if top != bottom {
			for i := right - 1; i >= left; i-- {
				if tmp[bottom][i+1] == 'Z' {
					tmp[bottom][i] = 'A'
					continue
				}
				tmp[bottom][i] = tmp[bottom][i+1] + 1
			}
		}
		if left != right {
			for  i := bottom - 1; i > top; i-- {
				if tmp[i+1][left] == 'Z' {
					tmp[i][left] = 'A'
					continue
				}
				tmp[i][left] = tmp[i+1][left] + 1
			}
		}
		left++
		right--
		top++
		bottom--
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res[i][j] = string(tmp[i][j])
		}
	}
	return res
}
