package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	res := []int{}
	i := 0
	for i < 1000 {
		inputReader := bufio.NewReader(os.Stdin)
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		input = input[:len(input)-1]
		if len(input) == 0 {
			break
		}
		tmp := strings.Split(input, " ")
		a, err := strconv.Atoi(tmp[0])
		b, err := strconv.Atoi(tmp[1])
		k, err := strconv.Atoi(tmp[2])
		v, err := strconv.Atoi(tmp[3])
		output := process_huowu(a, b, k, v)
		res = append(res, output)
		i++
	}
	for _, v := range res {
		fmt.Println(v)
	}
}

func process_huowu(a, b, k, v int) int {
	ge_1 := k-1
	max_1 := k * v
	jige := b / ge_1
	tmp_1 := a / max_1
	tmp_2 := a % max_1
	if tmp_2 > 0 {
		if tmp_1 < jige {
			tmp_1++
		}
	}
	if tmp_1 <= jige {
		return tmp_1
	}
	x_1 := jige * max_1
	y_1 := a - x_1
	tmp_3 := y_1 / v
	tmp_4 := y_1 % v
	if tmp_4 > 0 {
		tmp_3++
	}
	return tmp_3 + jige
}
