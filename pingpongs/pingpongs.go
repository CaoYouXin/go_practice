package main

import (
	"fmt"
)

func ping(ping chan<- string, msg string) {
	ping <- msg
}

func pong(ping <-chan string, pong chan<- string) {
	pong <- <-ping
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	go ping(pings, "whatever")
	go pong(pings, pongs)
	fmt.Println(<-pongs)
}
