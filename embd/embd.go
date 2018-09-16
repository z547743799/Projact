package main

import (
	"time"

	"github.com/kidoman/embd"
)

func main() {
	for {
		embd.LEDToggle("LED0")
		time.Sleep(200 * time.Millisecond)
	}
}
