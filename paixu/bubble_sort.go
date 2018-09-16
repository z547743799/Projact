package main

import "fmt"

func main() {
	ss := []int{23, 34, 1, 5, 1, 112}
	bubble(ss)
	fmt.Println(ss)
}
func bubble(s []int) {
	for i := 0; i < len(s); i++ {
		for l := i; l < len(s)-1; l++ {
			if s[i] > s[l+1] {
				s[i] = s[i] ^ s[l+1]
				s[l+1] = s[i] ^ s[l+1]
				s[i] = s[i] ^ s[l+1]
			}
		}
	}
}
