package main

import (
	"fmt"
)

func main() {
	msgs := make(chan string, 1)
	go func() {
		msgs <- "one"
		msgs <- "two"
		close(msgs)
	}()

	for msg := range msgs {
		fmt.Println(msg)
	}
}
