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

func readResult(results chan int) {
	v, ok := <-results
	for ok {
		fmt.Println(v)
	}

	close(results)
}

func main() {
	pool := make(chan int, 100)
	results := make(chan int, 100)

	//Create a worker to read result
	go readResult(results)

	//Create 3 worker look like 3 goroutine
	for i := 1; i <= 3; i++ {
		go worker(i, pool, results)
	}

	//Create 10 jobs
	for j := 1; j <= 10; j++ {
		pool <- j
	}

	close(pool)
}
