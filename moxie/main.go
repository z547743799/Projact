package main

import "fmt"

var Section = []int{12, 53, 65, 34}

func maopao(s []int) {
	for i := 0; i < len(s); i++ {
		for l := i; l < len(s)-1; l++ {
			if s[l] > s[l+1] {
				s[l], s[l+1] = s[l+1], s[l]
			}
		}
		fmt.Println(s)
	}
}
func main() {
	fmt.Println(Section)
	maopao(Section)
}
