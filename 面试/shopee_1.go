package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	_, _ = fmt.Scan(&s)
	res := process_reverse(s)
	fmt.Printf("%s", res)
}

func process_reverse(s string) string {
	str := strings.Split(s, "")
	n := len(str)
	if n == 0 {
		return ""
	}
	i := 0
	j := n - 1
	for i < j {
		if !((str[i] >= "a" && str[i] <= "z") || (str[i] >= "A" && str[i] <= "Z")) {
			i++
			continue
		}
		if !((str[j] >= "a" && str[j] <= "z") || (str[j] >= "A" && str[j] <= "Z")) {
			j--
			continue
		}
		str[i], str[j] = str[j], str[i]
		i++
		j--
	}
	return strings.Join(str, "")
}
