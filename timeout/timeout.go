package main

import (
	"fmt"
	"time"
)

func main() {
	chan1 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		chan1 <- "aha"
	}()
	select {
	case msg1 := <-chan1:
		fmt.Println(msg1)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1s")
	}

	chan2 := make(chan string)
	go func() {
		time.Sleep(2 * time.Second)
		chan2 <- "aha"
	}()
	select {
	case msg2 := <-chan2:
		fmt.Println(msg2)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 3s")
	}
}
