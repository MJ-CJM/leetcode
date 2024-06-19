package main

import "fmt"

func main() {
	exit := make(chan struct{})
	ch1, ch2 := make(chan struct{}), make(chan struct{})
	defer func() {
		close(exit)
		close(ch1)
		close(ch2)
	}()

	go func() {
		for i := 0; i <= 100; i += 2 {
			<-ch1
			fmt.Println(i)
			ch2 <- struct{}{}
		}
		exit <- struct{}{}
	}()

	go func() {
		for i := 1; i <= 101; i += 2 {
			<-ch2
			fmt.Println(i)
			ch1 <- struct{}{}
		}
		exit <- struct{}{}
	}()
	ch1 <- struct{}{}
	<-exit
}

