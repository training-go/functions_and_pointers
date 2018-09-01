package main

import (
	"fmt"
)

type person struct{
	age uint32
	name string
}

func (p person) SetName(name string) {
    p.name = name
}

func (p *person) SetAge(age uint32) {
    p.age = age
}

func main() {
	ivan := person{age: 29, name: "Ivan"}
	ivan.SetName("Petar")
	ivan.SetAge(20)
	fmt.Println("Hello, playground", ivan)
}
