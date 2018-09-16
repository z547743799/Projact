package main

import (
	"fmt"
)

	a = make(chan bool)
func main() {

	go test1()
	go test2()
	<-a
}
func test1() {
	for i := 1; i <= 10000; i++ {
		fmt.Println("tsrv", i)
	}
}
func test2() {
	for j := 1; j <= 1000; j++ {
		fmt.Println("234")
	}
	a<-ture
	close(a)


}
