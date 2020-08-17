package main

import "fmt"

//// 超出栈的空间了
//func myPow(x float64, n int) float64 {
//	output := new(float64)
//	*output = 1.0
//	if n < 0{
//		x = 1.0 / x
//		n = -n
//	}
//	*output = generate_new(x, n, output)
//	return *output
//}
//
//func generate_new(x float64, n int, output *float64) float64 {
//	// terminator
//	if n == 0{
//		return *output
//	}
//	// process logic
//	*output = *output * x
//	// drill down
//	return generate_new(x, n-1, output)
//	// reverse states
//}

// 暴力法求解，超出时间限制
//func myPow(x float64, n int) float64 {
//	result := 1.0
//	if n < 0{
//		x = 1.0 / x
//		n = -n
//	}
//	for n > 0{
//		result *= x
//		n -= 1
//	}
//	return result
//}

// 分治法
func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1.0 / x
		n = -n
	}
	result := fastcount(x, n)
	return result
}

func fastcount(f float64, i int) float64 {
	// terminator
	if i == 0{
		return 1.0
	}
	// process(split big problem)
	// drill down
	half := fastcount(f, i/2)
	// merge sub problems
	if i % 2 == 1{
		return half * half * f
	}else{
		return half * half
	}
}


func main(){
	x := 2.0
	//n1 := 10
	n2 := 10
	//fmt.Println(myPow(x, n1))
	fmt.Println(myPow(x, n2))
}