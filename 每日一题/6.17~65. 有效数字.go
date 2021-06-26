// -*- coding:utf-8 -*-
// @Time : 2021/6/17 11:27 下午
// @Author: MJ-CJM
// @File : leetcode/6.17~65. 有效数字
package main

import "strings"

// 遍历解法
func isNumber_travel(s string) bool {
	if len(s) == 0 {
		return false
	}
	isNum := false
	isDot := false
	ise_or_E := false
	for i := range s {
		if s[i] >= '0' && s[i] <= '9' {
			isNum = true
		} else if s[i] == '.' {
			if isDot || ise_or_E {
				return false
			}
			isDot = true
		} else if s[i] == 'e' || s[i] == 'E' {
			if !isNum || ise_or_E {
				return false
			}
			ise_or_E = true
			isNum = false
		} else if s[i] == '-' || s[i] == '+' {
			if i != 0 && s[i-1] != 'e' && s[i-1] != 'E' {
				return false
			}
		} else {
			return false
		}
	}
	return isNum
}

// 编译器扫描解法
func isNumber_2(s string) bool {
	pos := 0
	ls := len(s)
	n := byte(0)
	for pos < ls && s[pos] == '.' {
		pos++
	}
	// +-
	if pos < ls && (s[pos] == '+' || s[pos] == '-') {
		pos++
	}
	// [0-9]
	if p := scanNumber(s, pos, ls); p > pos {
		n |= 1
		pos = p
	}
	// .
	if pos < ls && s[pos] == '.' {
		pos++
		// [0-9]
		if p := scanNumber(s, pos, ls); p > pos {
			pos = p
			n |= 1 << 1
		}
	}
	// e
	if pos < ls && (s[pos] == 'e' || s[pos] == 'E') {
		pos++
		if pos < ls && (s[pos] == '+' || s[pos] == '-') {
			pos++
		}
		// [0-9]
		if p := scanNumber(s, pos, ls); pos == p {
			return false
		} else {
			pos = p
		}
	}
	for pos < ls && s[pos] == ' '{
		pos++
	}
	return n > 0 && pos ==ls
}

func scanNumber(s string, pos, end int) int {
	for pos < end && s[pos] >= '0' && s[pos] <= '9' {
		pos++
	}
	return pos
}


// 状态机官方解法
type State int
type CharType int

const (
	STATE_INITIAL State = iota
	STATE_INT_SIGN
	STATE_INTEGER
	STATE_POINT
	STATE_POINT_WITHOUT_INT
	STATE_FRACTION
	STATE_EXP
	STATE_EXP_SIGN
	STATE_EXP_NUMBER
	STATE_END
)

const (
	CHAR_NUMBER CharType = iota
	CHAR_EXP
	CHAR_POINT
	CHAR_SIGN
	CHAR_ILLEGAL
)

func toCharType(ch byte) CharType {
	switch ch {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return CHAR_NUMBER
	case 'e', 'E':
		return CHAR_EXP
	case '.':
		return CHAR_POINT
	case '+', '-':
		return CHAR_SIGN
	default:
		return CHAR_ILLEGAL
	}
}

func isNumber_1(s string) bool {
	transfer := map[State]map[CharType]State{
		STATE_INITIAL: map[CharType]State{
			CHAR_NUMBER: STATE_INTEGER,
			CHAR_POINT:  STATE_POINT_WITHOUT_INT,
			CHAR_SIGN:   STATE_INT_SIGN,
		},
		STATE_INT_SIGN: map[CharType]State{
			CHAR_NUMBER: STATE_INTEGER,
			CHAR_POINT:  STATE_POINT_WITHOUT_INT,
		},
		STATE_INTEGER: map[CharType]State{
			CHAR_NUMBER: STATE_INTEGER,
			CHAR_EXP:    STATE_EXP,
			CHAR_POINT:  STATE_POINT,
		},
		STATE_POINT: map[CharType]State{
			CHAR_NUMBER: STATE_FRACTION,
			CHAR_EXP:    STATE_EXP,
		},
		STATE_POINT_WITHOUT_INT: map[CharType]State{
			CHAR_NUMBER: STATE_FRACTION,
		},
		STATE_FRACTION: map[CharType]State{
			CHAR_NUMBER: STATE_FRACTION,
			CHAR_EXP:    STATE_EXP,
		},
		STATE_EXP: map[CharType]State{
			CHAR_NUMBER: STATE_EXP_NUMBER,
			CHAR_SIGN:   STATE_EXP_SIGN,
		},
		STATE_EXP_SIGN: map[CharType]State{
			CHAR_NUMBER: STATE_EXP_NUMBER,
		},
		STATE_EXP_NUMBER: map[CharType]State{
			CHAR_NUMBER: STATE_EXP_NUMBER,
		},
	}
	state := STATE_INITIAL
	for i := 0; i < len(s); i++ {
		typ := toCharType(s[i])
		if _, ok := transfer[state][typ]; !ok {
			return false
		} else {
			state = transfer[state][typ]
		}
	}
	return state == STATE_INTEGER || state == STATE_POINT || state == STATE_FRACTION || state == STATE_EXP_NUMBER || state == STATE_END
}

// 状态机
var (
	blank  = 0 // 空格
	digit1 = 1 // 数字(0-9) 无前缀
	sign1  = 2 // +/- 无e前缀
	point  = 4 // '.'
	digit2 = 5 // 数字(0-9) 有符号前缀
	e      = 6 // 'e'
	sign2  = 7 // +/- 有e前缀
	digit3 = 8 // 数字(0-9) 有e前缀
)

func isNumber(s string) bool {
	s = strings.TrimRight(s, " ")
	dfa := [][]int{
		[]int{blank, digit1, sign1, point, -1},
		[]int{-1, digit1, -1, digit2, e},
		[]int{-1, digit1, -1, point, -1},
		[]int{-1, digit2, -1, -1, e},
		[]int{-1, digit2, -1, -1, -1},
		[]int{-1, digit2, -1, -1, e},
		[]int{-1, digit3, sign2, -1, -1},
		[]int{-1, digit3, -1, -1, -1},
		[]int{-1, digit3, -1, -1, -1},
	}

	state := 0 // blank start
	for i := 0; i < len(s); i++ {
		var newState int
		switch s[i] {
		case ' ':
			newState = 0
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			newState = 1
		case '+', '-':
			newState = 2
		case '.':
			newState = 3
		case 'e':
			newState = 4
		default:
			return false
		}
		state = dfa[state][newState]
		if state == -1 {
			return false
		}
	}
	return state == digit1 || state == digit2 || state == digit3
}
