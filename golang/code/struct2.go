package main

import "fmt"

// START OMIT
type Person struct { // HL
	Name, Address string
}
type Employee struct {
	Person // HL
	Position string
}
func (e Employee) Describe() {
	fmt.Printf("%s, %s, works as a %s.\n", e.Name, e.Address, e.Position)
}
func main() {
	Employee{
		Person:   Person{"Jane", "456 Elm St"},
		Position: "Engineer",
	}.Describe()
}

// END OMIT
