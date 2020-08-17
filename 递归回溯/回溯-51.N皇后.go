package main

import "fmt"

func solveNQueens(n int) [][]string {
	result := [][]string{}
	pre_result := [][]int{}
	if n < 1{
		return result
	}
	cols := []int{}
	pie := []int{}
	na := []int{}
	row := 0
	iterm := []int{}
	dfs_process(n, row, iterm, &cols, &pie, &na, &pre_result)
	result = printQ(pre_result, n)
	return result
}

func dfs_process(n int, row int, iterm []int, cols *[]int, pie *[]int, na *[]int, pre_result *[][]int) {
	// terminator
	if row >= n{
		tmp := make([]int, n)
		for i, v := range iterm{
			tmp[i] = v
		}
		*pre_result = append(*pre_result, tmp)
		return
	}
	// process & drill down
	for col := 0; col < n; col ++{
		if is_in(col, cols) || is_in(row+col, pie) || is_in(row-col, na){
			continue
		}
		*cols = append(*cols, col)
		*pie = append(*pie, row+col)
		*na = append(*na, row-col)
		iterm = append(iterm, col)
		dfs_process(n, row+1, iterm, cols, pie, na, pre_result)
		// reverse states
		*cols = (*cols)[:len(*cols)-1]
		*pie = (*pie)[:len(*pie)-1]
		*na = (*na)[:len(*na)-1]
		iterm = iterm[:len(iterm)-1]
	}
}

func is_in(x int, l *[]int) bool{
	for _, v := range *l{
		if v == x{
			return true
		}
	}
	return false
}

// 打印q
func printQ(ints [][]int, n int) [][]string {
	result := [][]string{}
	for _, v := range ints{
		tmp := []string{}
		for _, q := range v{
			tmp2 := ""
			for j := 0; j < n; j++{
				if j == q {
					tmp2 += "Q"
					continue
				}
				tmp2 += "."
			}
			tmp = append(tmp, tmp2)
		}
		result = append(result, tmp)
	}
	return result
}


func main(){
	n := 6
	fmt.Println(solveNQueens(n))
}