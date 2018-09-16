package main

import "fmt"

func test(s *[]int) {
}
func main() {
	s := []int{2, 3, 4}
	test(&s)
	fmt.Println(s)

}
