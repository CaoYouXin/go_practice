package main

import (
	"fmt"
	"sync"
	"time"
)

var status = make(map[string]int)
var mutex = &sync.Mutex{}

func ping(ping chan<- string, pong <-chan string) {
	for {
		ping <- <-pong
		mutex.Lock()
		status["ping"]++
		mutex.Unlock()
	}
}

func pong(ping <-chan string, pong chan<- string) {
	for {
		pong <- <-ping
		mutex.Lock()
		status["pong"]++
		mutex.Unlock()
	}
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	pings <- "whatever"
	go ping(pings, pongs)
	go pong(pings, pongs)
	time.Sleep(time.Second)
	mutex.Lock()
	fmt.Println("status within 1 second:", status)
	mutex.Unlock()
}
