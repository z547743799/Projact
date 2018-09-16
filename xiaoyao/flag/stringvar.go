package main

import (
	"flag"
	"fmt"
)

var StringVar string

func main() {
	flag.StringVar(&StringVar, "dev", "1999", "consu")
	flag.Parse()
	fmt.Println(StringVar)
}
