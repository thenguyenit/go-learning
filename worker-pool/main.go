package main

import (
	"fmt"
	"time"
)

func worker(id int, pool <-chan int, results chan<- int) {
	for job := range pool {
		fmt.Println("Worker", id, " started job ", job)
		time.Sleep(time.Microsecond * 2)
		fmt.Println("Worker", id, " finished ", job)
		results <- job * 2
	}
}

func main() {
	const numJobs = 5
	pool := make(chan int, numJobs)
	results := make(chan int, numJobs)

	//Create 3 worker look like 3 goroutine
	for i := 1; i <= 3; i++ {
		go worker(i, pool, results)
	}

	//Create 10 jobs
	for j := 1; j <= numJobs; j++ {
		pool <- j
	}
	close(pool)

	for a := 1; a <= numJobs; a++ {
		fmt.Println("Result", <-results)
	}
}
