package main

import "fmt"

type example struct {
	a string
	b slice
	c inter
}

type functtion func(int)

type channle <-chan string

type slice []example

type slice1 []struct{}

type slice2 []interface{}

type slice3 []int

type hmap map[string]int

type hmap1 map[channle]int

type hmap2 map[int]func(int)

type hmap3 map[inter]slice

type hmap4 map[struct{}]slice

//map中的key 必须支持==运算符
//type hmap5 map[func(int)]slice
//type hmap5 map[functtion]slice
//type 定义的struct 类型无法放入map 的key值中
//type hmap5 map[example]slice

type inter interface {
	abba()
}

func main() {
	a := map[string]int{"go": 1}

	abba(nil, nil, a, nil, nil)
}

func abba(a functtion, b channle, c hmap, d slice, e inter) {
	fmt.Println(c["go"])

}
