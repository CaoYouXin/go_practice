package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key   int
	value int
	suc   chan bool
}

func main() {
	var readOps uint64
	var writeOps uint64

	reads := make(chan *readOp)
	writes := make(chan *writeOp)
	logs := make(chan bool)

	go func() {
		data := make(map[int]int)

		for {
			select {
			case read := <-reads:
				read.resp <- data[read.key]
			case write := <-writes:
				data[write.key] = write.value
				write.suc <- true
			case <-logs:
				fmt.Println("data:", data)
			}
		}
	}()

	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := &readOp{
					key:  rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &writeOp{
					key:   rand.Intn(5),
					value: rand.Intn(100),
					suc:   make(chan bool),
				}
				writes <- write
				if suc := <-write.suc; suc {
					atomic.AddUint64(&writeOps, 1)
					time.Sleep(time.Millisecond)
				}
			}
		}()
	}

	time.Sleep(time.Second)

	fmt.Println("read", atomic.LoadUint64(&readOps), "times")
	fmt.Println("write", atomic.LoadUint64(&writeOps), "times")
	logs <- true

	fmt.Scanln()
	fmt.Println("program end")
}
