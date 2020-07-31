package main

import "fmt"

type Employee struct {
	firstName, lastName string
}

func (e Employee) getFullName() string {
	return e.firstName + " " + e.lastName
}

func (e *Employee) changeFirstName(firstNameParam string) {
	e.firstName = firstNameParam
}

func main() {

	emp := &Employee{"Bala", "Krishna"}
	fmt.Println("My Full Name is : ", emp.getFullName())

	//Both statement is correct.
	// But it is not necessary to call method using pointer dereferencing variable
	//Because go will automatically convert variable in to pointer dereferencing syntax

	//(*emp).changeFirstName("balaji")
	emp.changeFirstName("balaji")
	fmt.Println("My Full Name After change is : ", emp.getFullName())
}
