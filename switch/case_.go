package main

import "fmt"

var s int

func main() {

	fmt.Scanln(s)
	s := 4

	switch {
	case 3 == s:
		fmt.Println("sdf")
	case 1 == s:
		fmt.Println("sdf")
	case 4 == s:
		fmt.Println("dsf")
	case s == 4:
		fmt.Println("dsf")
	}
}
