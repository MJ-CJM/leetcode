package main

import (
	"fmt"
	"math"
)

func main() {
	n, m := 0, 0
	_, _ = fmt.Scan(&n, &m)
	input := make([][]int, n)
	for i := 0; i < n; i++ {
		input[i] = make([]int, m)
	}
	for i := 0; i < n * m; i++ {
		x := 0
		_, _ = fmt.Scan(&x)
		input[i/m][i%m] = x
	}
	result := process_length(input)
	fmt.Printf("%d", result)
}

func process_length(input [][]int) int {
	n := len(input)
	m := len(input[0])

	max := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			result := 0
			max_length := math.MaxInt32
			k := 0
			back_length(input, max_length, i, j, k, &result)
			if result > max {
				max = result
			}
		}
	}
	return max
}

// [i-1,j], [i+1,j], [i, j-1], [i, j+1]
func back_length(input [][]int, example int,i int, j int, k int, resut *int){
	// terminator
	if i < 0 || j < 0 || i == len(input) || j == len(input[0]) {
		if k > *resut {
			tmp := k
			*resut = tmp
		}
		return
	}
	// process && drill down
	if input[i][j] < example {
		example = input[i][j]
		k++
		back_length(input, example, i-1, j, k, resut)
		back_length(input, example, i+1, j, k, resut)
		back_length(input, example, i, j-1, k, resut)
		back_length(input, example, i, j+1, k, resut)
	}
}
