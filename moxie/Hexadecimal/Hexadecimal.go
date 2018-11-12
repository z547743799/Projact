package main

import "fmt"

var a int

func main() {
	a = 123
	//十进制
	fmt.Printf("%d\n", a)
	//十六进制
	fmt.Printf("%x\n", a)
	//二进制
	fmt.Printf("%b\n", a)
	//八进制
	fmt.Printf("%o\n", a)
	//
	fmt.Printf("%T\n", a)

	fmt.Printf("%v\n", a)
}
