package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	r, err := os.Open("/home/lzw/f8e1f694a4c27d1e9f67eebb12d5ad6edcc4386e.jpg")
	if err != nil {
		panic(err)
	}
	defer r.Close()
	s := make([]byte, 0)

	b := make([]byte, 8)

	for {
		n, err := r.Read(b)
		if err != nil && err != io.EOF {
			panic(err)
		}
		if 0 == n {
			break
		}
		s = append(s, b...)

	}
	fmt.Println(s)
	err = ioutil.WriteFile("test.png", s, 0644)
	fmt.Println(err)
}
