package main

import "fmt"

type god struct {
	school string
}

func main() {
	abd := make([]*god, 0)
	as := god{"sdf"}
	abd = append(abd, as)
	fmt.Println(abd)
}
