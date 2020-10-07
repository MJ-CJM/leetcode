package main

import "fmt"

var luomap_2 = map[string]int{
	"I":  1,
	"IV": 4,
	"V":  5,
	"IX": 9,
	"X":  10,
	"XL": 40,
	"L":  50,
	"XC": 90,
	"C":  100,
	"CD": 400,
	"D":  500,
	"CM": 900,
	"M":  1000,
}

func romanToInt(s string) int {
	n := len(s)
	res := 0
	for i := 0; i < n; i++ {
		if i == n-1 {
			res += luomap_2[string(s[i])]
			continue
		}
		if luomap_2[string(s[i])] >= luomap_2[string(s[i+1])] {
			res += luomap_2[string(s[i])]
		}else{
			res += luomap_2[string(s[i+1])] - luomap_2[string(s[i])]
			i++
		}
	}
	return res
}

func main() {
	example := "MDCXCV"
	fmt.Println(romanToInt(example))
}
