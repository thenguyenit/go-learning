package main

import (
	"fmt"
	"time"
)

func square(c chan int) {
	fmt.Println("Cab", cap(c))
	fmt.Println("[square] Starting")
	n := <-c
	fmt.Println("[square] Read")
	c <- n * n
	fmt.Println("[square] Send")
}

func cube(c chan int) {
	fmt.Println("[cube] Starting")
	n := <-c
	fmt.Println("[cube] Read")
	c <- n * n * n
	fmt.Println("[cube] Send")
}

func main() {

	squareChan := make(chan int)
	cubeChan := make(chan int)

	fmt.Println("Cab", cap(squareChan))

	go square(squareChan)
	go cube(cubeChan)

	testNum := 3

	time.Sleep(2 * time.Second)
	fmt.Println("Sending to squareChan")
	squareChan <- testNum
	fmt.Println("Cab", cap(squareChan))
	fmt.Println("Size", len(squareChan))
	time.Sleep(2 * time.Second)

	fmt.Println("Sending to cubeChan")
	cubeChan <- testNum
	time.Sleep(2 * time.Second)

	squareVal, cubeVal := <-squareChan, <-cubeChan

	fmt.Printf("Square Val: %d\n", squareVal)
	fmt.Printf("Cube Val: %d\n", cubeVal)

}
