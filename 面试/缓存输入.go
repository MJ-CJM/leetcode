package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 声明并初始化带缓冲的读取器。
	// 准备从标准输入读取内容。
	//inputReader := bufio.NewReader(os.Stdin)
	//fmt.Println("Please input your name:")
	//// 以 \n 为分隔符读取一段内容。
	//input, err := inputReader.ReadString('\n')
	//if err != nil {
	//	fmt.Printf("Found an error : %s\n", err)
	//} else {
	//	// 对 input 进行切片操作，去掉内容中最后一个字节 \n 。
	//	input = input[:len(input)-1]
	//	fmt.Printf("Hello, %s!\n", input)
	//}
	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	input = input[:len(input)-1]
	fmt.Println(strings.Split(input, " "))
}
