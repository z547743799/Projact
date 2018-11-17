package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		分支语句：if，switch，select
		select 语句类似于 switch 语句，
			但是select会随机执行一个可运行的case。
			如果没有case可运行，它将阻塞，直到有case可运行。
			close()线程后通道仍可以传数据但是会报错误信息
	*/
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- 200
	}()
	go func() {
		time.Sleep(3 * time.Second)
		ch1 <- 100
	}()

	select {
	case num1 := <-ch1:
		fmt.Println("ch1中取数据。。", num1)
	case num2 := <-ch2:
		fmt.Println("ch2中取数据。。", num2)

	}

}