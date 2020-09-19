package main

import (
	"fmt"
	"math"
)

func main() {
	var x0 float64
	var y0 float64
	_, _ = fmt.Scan(&x0, &y0)
	var n int
	_, _ = fmt.Scan(&n)
	example := [][]float64{}
	for i := 0; i < n; i++ {
		var x1 float64
		var y1 float64
		_, _ = fmt.Scan(&x1, &y1)
		example = append(example, []float64{x1, y1})
	}
	res := process_juli(x0, y0, example)
	fmt.Printf("%.3f", res)
}

func process_juli(x0 float64, y0 float64, example [][]float64) float64 {
	var res float64
	res = math.MaxFloat64
	n := len(example)
	for i := 0; i < n-1; i++ {
		for j := i+1; j < n; j++ {
			x1 := example[i][0]
			y1 := example[i][1]
			x2 := example[j][0]
			y2 := example[j][1]
			tmp := count_juli(x0, y0, x1, y1, x2, y2)
			if tmp < res {
				res = tmp
			}
		}
	}
	return res
}

func count_juli(x0 float64, y0 float64, x1 float64, y1 float64, x2 float64, y2 float64) float64 {
	a := math.Abs((y2-y1)*x0 - (x2-x1)*y0 + y1*(x2-x1)-x1*(y2-y1))
	b := math.Sqrt((y2-y1)*(y2-y1)+(x2-x1)*(x2-x1))
	return a/b
}

