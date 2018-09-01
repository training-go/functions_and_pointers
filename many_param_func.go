package main

import (
	"fmt"
)

func sum(nums... int) int{
	total := 0
	for _, v := range nums {
 		total += v
	}
	return total 
}

func main() {
	fmt.Println(sum(2, 2, 6, 5,))
}
