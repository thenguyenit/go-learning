package main

import (
	"fmt"
	"sync"
)

var i int
var j int

func incWithMutex(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	j = j + 1
	m.Unlock()
	wg.Done()
}
func inc(wg *sync.WaitGroup) {
	i = i + 1
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go inc(&wg)
	}
	wg.Wait()
	fmt.Println(i)

	//With Mutex
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go incWithMutex(&wg, &m)
	}
	wg.Wait()
	fmt.Println(j)
}
