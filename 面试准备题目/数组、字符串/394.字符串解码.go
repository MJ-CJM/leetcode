package main

import (
	"strconv"
	"strings"
)

/*
使用栈：

使用两个栈：一个用于存储重复次数，另一个用于存储当前字符串。
遍历字符串，处理不同的字符：
如果是数字字符，解析出完整的数字并记录为重复次数。
如果是左括号 [，将当前字符串入栈，并重置当前字符串。
如果是右括号 ]，从栈中弹出一个重复次数，并将当前字符串重复该次数，然后与上一个字符串拼接。
如果是字母字符，直接添加到当前字符串中。
遍历字符串并处理：

初始化两个栈 countStack 和 stringStack，以及当前字符串 currentString 和当前重复次数 currentCount。
遍历字符串，根据不同的字符执行不同的操作。
*/
func decodeString(s string) string {
	countStack := []int{}
	stringStack := []string{}
	currentString := ""
	currentCount := 0

	for i := 0; i < len(s); i++ {
		char := s[i]

		if char >= '0' && char <= '9' {
			// 处理数字，可能是多位数
			numStart := i
			for i < len(s) && s[i] >= '0' && s[i] <= '9' {
				i++
			}
			currentCount, _ = strconv.Atoi(s[numStart:i])
			i-- // 调整 i，因为外层 for 循环还会增加 1
		} else if char == '[' {
			// 将当前字符串和重复次数入栈
			countStack = append(countStack, currentCount)
			stringStack = append(stringStack, currentString)
			// 重置当前字符串和重复次数
			currentString = ""
			currentCount = 0
		} else if char == ']' {
			// 处理结束的字符串
			repeatTimes := countStack[len(countStack)-1]
			countStack = countStack[:len(countStack)-1]
			tempString := stringStack[len(stringStack)-1]
			stringStack = stringStack[:len(stringStack)-1]
			currentString = tempString + strings.Repeat(currentString, repeatTimes)
		} else {
			// 处理普通字符
			currentString += string(char)
		}
	}

	return currentString
}
