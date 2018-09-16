package main

import "fmt"

func main() {
	a := 1
	b := 0
	c := 0
	for i := 0; i < 12; i++ {
		c = a + b
		a = b
		b = c
		fmt.Println(c)
	}
}
