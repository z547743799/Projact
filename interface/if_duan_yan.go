package main

import "fmt"

type sss interface{}

type as struct {
	aa string
}

func main() {
	var s sss
	s = "df"
	if f, o1 := s.(string); o1 {
		fmt.Println(f)
	}

}
