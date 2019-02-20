package main

import "fmt"

type Person struct {
	name string
	age int
}

func (v Person) SetName(newName string) string{
	v.name=newName
	return v.name
}

func (v *Person) SetAge(newAge int) int{
	v.age=newAge
    return v.age
}

func main()  {

	s := make([]int,5)
	fmt.Print(s)

	//v:=Person{"Nikola",21}
	//fmt.Println(v)
	//v.SetName("Pesho")
	//v.SetAge(19)
	//fmt.Println(v)
}