package main

import "fmt"

// 封闭模式
func main() {
	exit := make(chan bool)
	ch1, ch2 := make(chan bool), make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- true
			fmt.Println("A", i)
			//在ch1和ch2之间是阻塞独占的
			<-ch2
		}
		exit <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			<-ch1
			fmt.Println("B", i)
			ch2 <- true
		}
	}()
	<-exit
	huanchong()
}

// 缓冲模式
func huanchong() {
	exit := make(chan bool)
	ch1, ch2 := make(chan bool, 1), make(chan bool, 1)
	ch1 <- true // 生产（选择一个启动项）

	go func() {
		for i := 1; i < 10; i++ {
			if ok := <-ch1; ok {
				fmt.Println("A", 2*i-1)
				ch2 <- true
			}
		}
	}()

	go func() {
		defer func() { close(exit) }()
		for i := 1; i < 10; i++ {
			if ok := <-ch2; ok {
				fmt.Println("B", 2*i)
				ch1 <- true
			}
		}
	}()
	<-exit
}
