package main

// Person represents the function found in the source code
type Person struct {
	Name string // Name is the name of the person we are reviewing
	Age  int    // Age is the age of the person we are reviewing
}

// SetName sets the name of Person
func (v Person) SetName(newName string) string {
	v.Name = newName
	return v.Name
}

// SetAge sets the age of Person
func (v *Person) SetAge(newAge int) int {
	v.Age = newAge
	return v.Age
}
