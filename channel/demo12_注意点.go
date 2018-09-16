package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		go 函数名()
	*/
	//for i:=1;i<=5;i++{
	//	go func(n int) {
	//		//fmt.Println("第",i,"个 goroutine。。")
	//		fmt.Println("第",n,"个 goroutine。。")
	//	}(i)
	//	//time.Sleep(1*time.Second)
	//}

	for i := 1; i <= 5; i++ {
		n := i //1
		go func() {
			fmt.Println("第", n, "个goroutine。。")
		}()
	}
	time.Sleep(3 * time.Second)
}
