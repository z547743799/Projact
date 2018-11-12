package main

import "fmt"

func main() {
	ch1 := make(chan int)
	done := make(chan bool)
	//生产者1
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("生产者1产生数据：", i) // 0
			ch1 <- i                    //
		}
		done <- true
	}()
	// 生产者2
	go func() {
		for i := 100; i < 200; i++ {
			ch1 <- i                    // 100
			fmt.Println("生产者2产生数据：", i) //100
		}
		done <- true
	}()

	//消费者
	go func() {
		for n := range ch1 { // 100
			fmt.Println("\t消费者消费数据：", n) // 0 1
		}
		done <- true
	}()

	//main goroutine
	<-done
	<-done
	close(ch1) //两个生产者生产数据后，关闭通道

	<-done
	fmt.Println("main..over....")
}
