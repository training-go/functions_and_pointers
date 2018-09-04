package main

import (
	"fmt"
)
// Person contains info for the user
type Person struct{
	age uint32
	name string
}

// SetName sets the name 
func (p Person) SetName(name string) {
    p.name = name
}

func (p *Person) SetAge(age uint32) {
    p.age = age
}

func main() {
	ivan := Person{age: 29, name: "Ivan"}
	ivan.SetName("Petar")
	ivan.SetAge(20)
	fmt.Println("Hello, playground", ivan)
}

