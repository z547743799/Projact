package main

import "fmt"

var ss = []int{23, 43, 2}

func main() {
	date := struct {
		sd []int
	}{
		sd: ss,
	}
	fmt.Println(date)
}
