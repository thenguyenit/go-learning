package main

import "fmt"

func main() {
	a := fibonacci(1)
	fmt.Println(a)

	a = fibonacci(5)
	fmt.Println(a)
}

func fibonacci(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonacci(i-1) + fibonacci(i-2)
}
