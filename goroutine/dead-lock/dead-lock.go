package main

import "fmt"

func greet(c chan string) {
	fmt.Println("Greet function start")
	// fmt.Print(<-c)
	// fmt.Print(<-c)
	close(c)
}

func main() {
	fmt.Println("Start")

	c := make(chan string)

	go greet(c)

	c <- "John"

	c <- "John1"

	// close(c)

	fmt.Println("Next to push string to Channe")

	fmt.Println("Finished")
}
