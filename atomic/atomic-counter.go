package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var ops uint64

	for i := 0; i < 100; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)
	fmt.Println(atomic.LoadUint64(&ops))
}