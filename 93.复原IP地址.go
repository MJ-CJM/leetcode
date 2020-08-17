package main

import (
	"fmt"
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	if len(s) < 4 || len(s) > 12 {
		return nil
	}
	result := []string{}
	iterm := []string{}
	level := 1
	backrack(s, iterm, level, &result)
	return result
}

func backrack(s string, iterm []string, level int, result *[]string) {
	// terminator
	if level == 5 {
		if len(s) == 0 {
			str := strings.Join(iterm, ".")
			*result = append(*result, str)
		}
		return
	}

	// process && drill down
	for j := 1; j <= 3; j++ {
		if j <= len(s){
			v, err := strconv.Atoi(s[:j])
			if err == nil && v <= 255{
				c1 := append(iterm, s[:j])
				str2 := s[j:]
				backrack(str2, c1, level+1, result)
			}
			if v == 0{
				break
			}
		}
	}
}

func main() {
	s := "25525511135"
	fmt.Println(restoreIpAddresses(s))
}
