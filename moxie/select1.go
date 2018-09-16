package main

import (
	"fmt"
	"sync"
)

var m sync.RWMutex

func main() {
	ch1 := make(chan int)
	down := make(chan bool)
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
			fmt.Printf("%d--\n", i)
		}
		close(ch1)
	}()
	go func() {
		s := 0
		for 99 >= s {
			s = <-ch1

			select {
			case <-ch1:
				fmt.Println(s)
			}

		}
		down <- true

	}()
	<-down
}
