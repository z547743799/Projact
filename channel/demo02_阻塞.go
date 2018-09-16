package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	done := make(chan bool) // 通道
	go func() {
		fmt.Println("子goroutine执行。。。")
		time.Sleep(3 * time.Second)
		data := <-ch1 // 从通道中读取数据
		fmt.Println("data：", data)
		done <- true
	}()
	// 向通道中写数据。。
	time.Sleep(5 * time.Second)
	ch1 <- 100

	<-done
	fmt.Println("main。。over")

}

/*
练习1：创建并启动一个子 goroutine，打印100个数字，要保证在main goroutine结束前结束。
	通道实现。
练习2：创建并启动两个子goroutine，一个打印100个数字，另一个打印100个字母，要保证在main goroutine结束前结束。
*/
