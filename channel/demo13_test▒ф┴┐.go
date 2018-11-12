package main

import "fmt"

func main() {
	i := 100
	fmt.Println("i-->", i)

	{
		i := 20
		fmt.Println("i------>", i)
	}
}
