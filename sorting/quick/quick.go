package main

import "fmt"

var numbers []int

func main() {
	numbers = []int{2, 1, 100, 89, 5, 3, 0, -4, 7, 8, 7}
	fmt.Println(numbers)
	numbers = quickSort(numbers)
	fmt.Println(numbers)
}

func quickSort(s []int) []int {

	if len(s) <= 1 {
		return s
	}

	pivot := len(s) / 2

	var lt, gt, md []int
	x := s[pivot]

	for _, v := range s {
		switch {
		case v < x:
			lt = append(lt, v)
		case v == x:
			md = append(md, v)
		case v > x:
			gt = append(gt, v)
		}
	}

	lt = quickSort(lt)
	gt = quickSort(gt)

	//Merge
	lt = append(lt, md...)
	lt = append(lt, gt...)

	return lt
}
