package main

import "fmt"

type test struct {
	id   int
	name string
}

var map1 = make(map[int]*test)

func main() {
	s := new(test)
	s.id = 1
	s.name = "123"
	a := test{id: 1, name: "23"}
	fmt.Println(s, a)
	map1[0] = s
	map1[1] = &a
	for i := 0; i < 2; i++ {
		fmt.Println(map1[i])
	}
}
