package main

import "fmt"

var values = make(map[string][]string)

func main() {

	values["sdf"] = []string{"dsf", "sdf", "sdf"}
	in := make([]interface{}, 3)
	in[0] = values
	fmt.Println(in[0])
}
