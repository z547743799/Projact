package main

import (
	"fmt"

	"github.com/edwingeng/wuid/redis"
)

func main() {

	// Setup
	g := wuid.NewWUID("default", nil)
	g.LoadH24FromRedis("127.0.0.1:6379", "", "wuid")

	// Generate
	fmt.Println(g.Next())
}
