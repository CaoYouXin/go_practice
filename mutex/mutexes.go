package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var readOps uint64
	var writeOps uint64

	data := make(map[int]int)
	mutex := &sync.Mutex{}

	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += data[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				mutex.Lock()
				data[key]++
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	fmt.Println("read", atomic.LoadUint64(&readOps), "times")
	fmt.Println("write", atomic.LoadUint64(&writeOps), "times")
	mutex.Lock()
	fmt.Println("data:", data)
	mutex.Unlock()
}
