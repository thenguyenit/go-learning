package main

import (
	"fmt"
	"time"
)

func worker(id int, pool <-chan int, results chan<- int) {
	//The loop for job := range pool receives values from the channel repeatedly until it is closed.
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

	//Create jobs and push them to pool
	for j := 1; j <= numJobs; j++ {
		pool <- j
	}
	//close a channel to indicate that no more values will be sent
	//Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.
	close(pool)

	//Read the result and block main routine until no more values from the results channel 
	for a := 1; a <= numJobs; a++ {
		fmt.Println("Result", <-results)
	}
}
