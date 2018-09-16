package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan bool, 1)
	go func() {
		defer close(ch1)
		for {
			time.Sleep(time.Second * 1)
			hor, minute, second := time.Now().Clock()
			fmt.Println(hor, ":", minute, ":", second)
			if time.Now().Second() == 50 {
				break
			}
		}
		ch1 <- true
	}()
	s := <-ch1
	fmt.Println(s)
}
