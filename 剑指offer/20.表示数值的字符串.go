package main

import (
	"regexp"
	"strings"
)

func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	// res,_ := regexp.MatchString("^(([\\+\\-]?[0-9]+(\\.[0-9]*)?)|([\\+\\-]?\\.?[0-9]+))(e[\\+\\-]?[0-9]+)?$", s)
	res,_ := regexp.MatchString("^(([\\+\\-]?[0-9]+(\\.[0-9]*)?)|([\\+\\-]?\\.?[0-9]+))((e|E)[\\+\\-]?[0-9]+)?$", s)

	return res
}

