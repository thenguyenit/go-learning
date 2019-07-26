package main

import "fmt"

func main() {
	numbers := [8]int{2, 1, 100, 89, 5, 3, 0, -4}
	nL := len(numbers)

	fmt.Println(numbers)

	for i, _ := range numbers {
		for j := 0; j < nL-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				tmp := numbers[j+1]
				numbers[j+1] = numbers[j]
				numbers[j] = tmp
			}
		}
	}
}
