package main

import "fmt"

type name struct {
	sww string
}

type shoocl struct {
	name
	sww string
}

func main() {
	var aa = shoocl{}
	aa.sww = "sdf"
	fmt.Println(aa.sww)
}
