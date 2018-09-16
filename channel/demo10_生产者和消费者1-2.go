package main

import "fmt"

func main() {
	/*
		一条goroutine，生产数据，理解为生产者
		两条goroutine，获取数据，理解为消费者
	*/
	ch1 := make(chan int)
	done := make(chan bool)
	//生产者
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i                   // 0 1 2 3
			fmt.Println("生产者产生数据：", i) // 0 1 2
		}
		close(ch1) //生产完毕后，关闭通道
	}()
	//消费者1
	go func() {
		for n := range ch1 {
			fmt.Println("\t消费者range1：", n)
		}
		done <- true
	}()
	//消费者2
	go func() {
		for n := range ch1 { // 0 1 2  3
			fmt.Println("\t消费者range2：", n) // 0 1 2 3
		}
		done <- true
	}()

	<-done
	<-done
	fmt.Println("main。。over。。。。。")
}
