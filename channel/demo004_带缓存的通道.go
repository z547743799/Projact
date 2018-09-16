package main

import (
	"fmt"
)

var ch6 chan int
var ch7 chan int

//var ch8 chan int
func main() {
	ch6 = make(chan int)
	ch7 = make(chan int)
	//ch8=make(chan int)
	go w()
	go r()
	<-ch7
	<-ch7

}
func w() {
	for i := 0; i <= 10; i++ {
		ch6 <- i
		fmt.Println("w写入通道", i)

	}
	close(ch6)
	ch7 <- 100

}
func r() {

	for {

		v, ok := <-ch6
		if !ok {
			fmt.Println("写完了", ok)
			break

		}
		fmt.Println("r读取数据", v)

	}
	ch7 <- 100

}
