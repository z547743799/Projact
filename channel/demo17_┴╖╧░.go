package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)
	done := make(chan bool)

	go func() {
	out:
		for {
			//time.Sleep(1*time.Second)
			select {
			case num1, ok := <-ch1:
				fmt.Println("case1执行。。", num1, ok)

			case num2, ok := <-ch2:
				fmt.Println("case2执行。。", num2, ok)
				if !ok {
					fmt.Println("case2结束。。")
					break out
				}
			}

		}
		fmt.Println("子goroutine中，结束for循环。。")
		done <- true
	}()

	ch1 <- 100
	ch2 <- "hello"

	close(ch1)
	close(ch2)

	<-done

}
