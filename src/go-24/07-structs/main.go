package main

import (
	"fmt"
)

type Employee struct {
	Name string
	Age  int
}

func NewEmployee(name string) Employee {
	//for setting default value for Age
	emp := Employee{
		Name: name,
		Age:  25,
	}

	return emp
}

func main() {

	// Initializing using short variable assignment
	emp := Employee{
		Name: "Subhransu",
		Age:  24,
	}

	fmt.Printf("%+v\n", emp)

	// -------------------------------------------
	//Using new
	emp2 := new(Employee)
	emp2.Name = "OOlala"
	emp2.Age = 25

	fmt.Printf("%+v\n", emp2)

	//Setting default value for Age using custom constructor
	fmt.Printf("%+v\n", NewEmployee("Maharana"))
}
