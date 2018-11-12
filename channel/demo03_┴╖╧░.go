package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
		channel,属于引用传递
	*/
	ch1 := make(chan bool)
	//fmt.Println(ch1)

	go printNum(ch1)
	go printLetter(ch1)
	<-ch1
	<-ch1

	fmt.Println("main...over.....")
}

func printNum(ch1 chan bool) {
	//fmt.Println("子goroutine，ch1，", ch1)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		fmt.Println("子goroutine1，打印i：", i)
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	}
	ch1 <- true
}

func printLetter(ch1 chan bool) {
	rand.Seed(time.Now().UnixNano())
	for j := 0; j < 100; j++ {
		fmt.Println("子goroutine2，打印字母j：", j)
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
	}
	ch1 <- true //程序结束，向通道中写true
}
