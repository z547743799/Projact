package main

func main() {
	/*
		用于参数传递：
	*/
	ch1 := make(chan int)   //双向，读，写
	ch2 := make(chan<- int) // 单向，只写，不能读
	ch3 := make(<-chan int) //单向，只读，不能写
	//ch1 <- 100
	//data :=<-ch1
	//ch2 <- 1000
	//	<-ch2 //invalid operation: <-ch2 (receive from send-only type chan<- int)
	//	<-ch3
	//	ch3 <- 100 //invalid operation: ch3 <- 100 (send to receive-only type <-chan int)

	fun1(ch2)
	fun1(ch1)
	fun2(ch3)
	fun2(ch1)
}

//该函数接收，只写的通道
func fun1(ch chan<- int) {
	// 函数内部，对于ch只能写数据，不能读数据
}

func fun2(ch <-chan int) {
	//函数内部，对于ch只能读数据，不能写数据
}
