package main

import (
	"testing"
)

func TestSumOfRandomSlice(t *testing.T) {
	if Sum(5, 9, 2, -5, 6) != 17 {
		t.Errorf("Sum was not corrent, want %d, got %d", 17, Sum(5, 9, 2, -5, 6))
	}
	if Sum(2, 1, -222) != -219 {
		t.Errorf("Sum was not corrent, want %d, got %d", -219, Sum(2, 1, -222))
	}
	s := []int{3, 6, 10, 4, -5, -9, 5, 77, 2, 5, 6, 21, 34, 52, -29}
	if Sum(s...) != 182 {
		t.Errorf("Sum was not corrent, want %d, got %d", 182, Sum(s...))
	}
}
