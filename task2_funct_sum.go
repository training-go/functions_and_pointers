package main

import (
	"fmt"
)

// Sumator takes two params and sets a result
type Sumator struct{
	a int
	b int
	result int
}

func (s *Sumator) sum(){
	s.result = s.a + s.b
}

func main() {
	new_sumator := Sumator{1,2,0}
	
	fmt.Println("Hello, playground", new_sumator.a)
}
