package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 3)
	done := make(chan bool, 1)

	go func() {
		for {
			time.Sleep(time.Second * 2)
			j, more := <-jobs
			if more {
				fmt.Println("work with", j)
			} else {
				done <- true
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}
