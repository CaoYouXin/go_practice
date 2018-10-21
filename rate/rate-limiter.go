package main

import (
	"fmt"
	"time"
)

func rateLimiter1() {
	requests := make(chan int, 5)
	for r := 0; r < 5; r++ {
		requests <- r
	}
	close(requests)

	timer := time.Tick(200 * time.Millisecond)
	for r := range requests {
		<-timer
		fmt.Println("handle request 1", r)
	}
}

func rateLimiter2() {
	requests := make(chan int, 5)
	for r := 0; r < 5; r++ {
		requests <- r
	}
	close(requests)

	timer := make(chan time.Time, 3)
	for t := 0; t < 3; t++ {
		timer <- time.Now()
	}
	go func() {
		for t := range time.Tick(1000 * time.Millisecond) {
			timer <- t
		}
	}()

	for r := range requests {
		<-timer
		fmt.Println("handle request 2", r)
	}
}

func main() {
	rateLimiter1()
	rateLimiter2()
}
