package main

import "fmt"

//Person represents the function found in the source code
type Person struct {
	Name string // name of the person we are reviewing
	Age  int    // age of the person we are reviewing
}

//set the age of person
func (v Person) SetName(newName string) string {
	v.Name = newName
	return v.Name
}

//set the name of person
func (v *Person) SetAge(newAge int) int {
	v.Age = newAge
	return v.Age
}

func main() {
	v := Person{"Nikola", 21}
	fmt.Println(v)
	v.SetName("Pesho")
	v.SetAge(19)
	fmt.Println(v)
}
