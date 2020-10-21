package main

func myPow(x float64, n int) float64 {
	if x == 0 {
		return x
	}
	if n == 0 {
		return 1.0
	}
	if n == 1 {
		return x
	}
	res := 1.0
	if n > 0 {
		res = myPow(x, n/2)
		res = res * res * myPow(x, n%2)
	}else{
		n = -n
		res = myPow(x, n/2)
		res = res * res * myPow(x, n%2)
		res = 1 / res
	}
	return res
}
