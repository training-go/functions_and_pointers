package main

// Sum returns the sum of an undefined number of numbers
func Sum(n ...int) int {
	s := 0
	for _, v := range n {
		s += v
	}
	return s
}
