package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	go sendData2(ch1)
	// for循环的for range形式可用于从通道接收值，直到它关闭为止。手动调用close()
	for v := range ch1 {
		fmt.Println("读取数据：", v)
	}
	fmt.Println("main..over.....")
}
func sendData2(ch1 chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		ch1 <- i
	}
	close(ch1) //通知对方，通道关闭
}
