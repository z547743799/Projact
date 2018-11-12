package main

import "fmt"

type nn struct {
	s string
}

func main() {
	n := make([]interface{}, 0)

	sn := new(nn)

	sn.s = "---"

	n = append(n, sn)
	fmt.Println(n)
}
