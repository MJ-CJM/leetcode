package main

import (
	"fmt"
)


func generateParenthesis(n int) []string {
	Output := new([]string)
	_generate(0, 0, n, "", Output)
	return *Output
}

func _generate(left int, right int, max int, s string, Output *[]string){
	// terminator
	if left == right && left ==  max{
		*Output = append(*Output, s)
		return
	}
	// process current logic
	// drill down
	s1 := s
	s2 := s
	if left < max{
		s1 = s + "("
		_generate(left+1, right, max, s1, Output)
	}
	if right < left{
		s2 = s + ")"
		_generate(left, right+1,  max, s2, Output)
	}

	// reverse states
}

func main(){
	n := 1
	output := generateParenthesis(n)
	fmt.Println(output)
}