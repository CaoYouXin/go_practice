package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "starts job", j)
		time.Sleep(500 * time.Millisecond)
		fmt.Println("worker", id, "finishes job", j)
		results <- j
	}
}

func main() {
	jobs := make(chan int)
	results := make(chan int)
	done := make(chan bool)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	go func() {
		for j := 0; j < 10; j++ {
			<-results
		}
		done <- true
	}()

	for j := 0; j < 10; j++ {
		jobs <- j
	}
	close(jobs)

	<-done
	fmt.Println("Program finished!")
}
