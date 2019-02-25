package main

import "fmt"

//Contains the details of the person
type Person struct {
	name string
	age  int
}

//set age of person
func (v Person) SetName(newName string) string {
	v.name = newName
	return v.name
}

//set name of person
func (v *Person) SetAge(newAge int) int {
	v.age = newAge
	return v.age
}

func main() {
	v := Person{"Nikola", 21}
	fmt.Println(v)
	v.SetName("Pesho")
	v.SetAge(19)
	fmt.Println(v)
}
