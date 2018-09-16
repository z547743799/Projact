package main

import "fmt"

func main() {
	/*
	练习1：启动一个子goroutine，向带缓存的通道中写出数据，
	另一个子goroutine，从该通道中读取数据。
	*/
	ch1 := make(chan int, 5)
	done := make(chan bool)
	go func() {
		//读
		for n := range ch1 {
			fmt.Println("读到的数据是：", n)
		}
		done <- true
	}()

	go func() {
		//写
		for i := 0; i < 100; i++ {
			ch1 <- i
			fmt.Println("\t写出数据：", i)
		}
		close(ch1)
	}()

	<-done
	fmt.Println("main。。over。。。。")
}
