package main

import "fmt"

func main() {
	chan1 := make(chan string, 2)
	chan2 := make(chan string, 2)

	chan1 <- "V1"
	chan1 <- "V2"
	// chan1 <- "V11"

	chan2 <- "V3"
	chan2 <- "V4"

	fmt.Println(<-chan1)
	fmt.Println(<-chan1)
	fmt.Println(<-chan2)
	fmt.Println(<-chan2)

	// select {
	// case res := <-chan1:
	// 	fmt.Println("Res from chan1", res)
	// 	fmt.Println(<-chan1)
	// case res := <-chan2:
	// 	fmt.Println("Res from chan2", res)
	// 	fmt.Println(<-chan2)
	// }

	fmt.Println("Closed")

}
