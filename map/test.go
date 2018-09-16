package main

import (
	"fmt"
)

type ss func(int)
type assr func(string) int

//从struct类型方法中选择实现方法
var rpc = map[string]struct {
	sz  ss
	ssz assr
}{
	"sdfsdf": {sz: sb},
	"sdf":    {ssz: sbb},
}

func sb(s int) {
	fmt.Println(s)
}
func sbb(s string) int {
	b := []byte(s)

	//a := strconv.Itoa(s)
	fmt.Println(b)

	return 1
}
func main() {
	ssb, ok := rpc["sdf"]
	if !ok {
		fmt.Println("dsf1")
	}
	ssb.ssz("sdf")

}
