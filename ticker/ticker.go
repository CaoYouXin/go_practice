package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(300 * time.Millisecond)
	go func() {
		for t := range ticker.C {
			fmt.Println("tick at", t)
		}
	}()
	time.Sleep(time.Second)
	ticker.Stop()
	fmt.Println("ticker has been stopped")
}
