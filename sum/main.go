package main

import "fmt"

// Sum returns the sum of an undefined number of numbers
func Sum(n ...int) int {
	s := 0
	for _, v := range n {
		s += v
	}
	return s
}

func main() {
	fmt.Println(Sum(5, 9, 2, -5, 6))
	fmt.Println(Sum(2, 1, -222))
	s := []int{3, 6, 10, 4, -5, -9, 5, 77, 2, 5, 6, 21, 34, 52, -29}
	fmt.Println(Sum(s...))
}
