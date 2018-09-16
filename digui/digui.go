package main

import (
	"Project/digui/test"
	"fmt"
)

const (
	jsonrpcSem = 1
)

type requestRecursive func(string)

type requestTransformation func(rune) []byte

func init() {
	fmt.Printf("sdf\n")

}

var spe = map[string]struct {
	Recursive      requestRecursive
	Transformation requestTransformation
}{
	"Abelive": {Recursive: test.Abelive},
	"belive":  {Transformation: test.Belive},
}

func main() {

	//递归打印目录
	//sfd, ok := spe["Abelive"]
	//sfd.Recursive("/home/lzw/") //印了目录
	fmt.Println("-------")
	//rune转型
	sdf, ok := spe["belive"]
	if ok {
		fmt.Println("sdfsdf")
	}
	l := sdf.Transformation('是') //[47]
	fmt.Println(l)
	fmt.Println("-------")
	//robot

}
