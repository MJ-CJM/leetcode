package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main() {
	var tmp string
	for i := 0; i < math.MaxInt32; i++ {
		tmp += "a"
	}
	fmt.Println("main size", unsafe.Sizeof(tmp))
}
