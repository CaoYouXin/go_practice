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
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 0; j < 10; j++ {
		jobs <- j
	}
	// close(jobs)

	for j := 0; j < 10; j++ {
		<-results
	}
}
