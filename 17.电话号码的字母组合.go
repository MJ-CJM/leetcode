package main

import "fmt"

func letterCombinations(digits string) []string {
	if digits == ""{
		return nil
	}
	result := []string{}
	number_map := make(map[string][]string)
	number_map["2"] = []string{"a", "b", "c"}
	number_map["3"] = []string{"d", "e", "f"}
	number_map["4"] = []string{"g", "h", "i"}
	number_map["5"] = []string{"j", "k", "l"}
	number_map["6"] = []string{"m", "n", "o"}
	number_map["7"] = []string{"p", "q", "r", "s"}
	number_map["8"] = []string{"t", "u", "v"}
	number_map["9"] = []string{"w", "x", "y", "z"}
	index := 0
	iterm := ""
	_help(digits, number_map, iterm, &result, index)
	return result
}

func _help(s string, number_map map[string][]string, iterm string, result *[]string, index int) {
	// terminator
	if index == len(s) {
		*result = append(*result, iterm)
		return
	}
	// process,split big problems
	numbers := number_map[string(s[index])]
	n := len(numbers)
	for i := 0; i < n; i++{
		// drill down
		_help(s, number_map, iterm + numbers[i], result, index + 1)
		// reverse state
	}
}

func main(){
	nums := "23"
	fmt.Println(letterCombinations(nums))
}