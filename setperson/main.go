package main

import "fmt"

type person struct {
	name string
	age int
}
//set age of person
func (v person) SetName(newName string) string{
	v.name=newName
	return v.name
}
//set name of person
func (v *person) SetAge(newAge int) int{
	v.age=newAge
    return v.age
}

func main()  {
	v:=person{"Nikola",21}
	fmt.Println(v)
	v.SetName("Pesho")
	v.SetAge(19)
	fmt.Println(v)
}