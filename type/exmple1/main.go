package main

import "fmt"

type example struct {
	s string
	a int
	c complex64
}

type functtion func(int)

type channle <-chan string

type hmap map[string]int

type slice []int

type inter interface{}

func main() {
	a := map[string]int{"go": 1}
	abba(nil, nil, a, nil, nil)
}

func abba(a functtion, b channle, c hmap, d slice, e inter) {
	fmt.Println(c["go"])

}
