package main

import "fmt"

func main() {
	/*
		Channel，通道，
			1.用于goroutine，传递消息的。
			2.通道，每个都有相关联的数据类型,
				nil chan，不能使用，类似于nil map，不能直接存储键值对
			3.使用通道传递数据：<-
				chan <- data,发送数据到通道。向通道中写数据
				data <- chan,从通道中获取数据。从通道中读数据
			4.阻塞：
				发送数据：chan <- data,阻塞的，直到另一条goroutine，读取数据来解除阻塞
				读取数据：data <- chan,也是阻塞的。直到另一条goroutine，写出数据解除阻塞。
			5.本身channel就是同步的，意味着同一时间，只能有一条goroutine来操作。
	*/
	var ch1 chan bool       //声明，没有创建
	fmt.Println(ch1)        //<nil>
	fmt.Printf("%T\n", ch1) //chan bool
	ch1 = make(chan bool)   //0xc042036060,是引用类型的数据
	fmt.Println(ch1)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("子goroutine中，i：", i)
		}
		// 循环结束后，向通道中写数据，表示要结束了。。
		ch1 <- true

		fmt.Println("结束。。")

	}()

	data := <-ch1 // 从ch1通道中读取数据
	fmt.Println("data-->", data)
	fmt.Println("main。。over。。。。")
}
