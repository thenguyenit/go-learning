package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Start")
	var waitgroup sync.WaitGroup
	waitgroup.Add(1)
	go func(waitgroup *sync.WaitGroup) {
		fmt.Println("Go routin")
		waitgroup.Done()
	}(&waitgroup)
	waitgroup.Wait()
	fmt.Println("End")
}
