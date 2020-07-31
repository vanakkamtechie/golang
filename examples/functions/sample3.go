package main

import "fmt"

type Address struct {
	place string
}

type Employee struct {
	firstName, lastName string
	address             Address
}

func (e Employee) getFullName() string {
	return e.firstName + " " + e.lastName
}

func (e *Employee) changeFirstName(firstNameParam string) {
	e.firstName = firstNameParam
}

func (e *Employee) changeAddress(place string) {
	e.address.place = place
}

func (e Employee) getAddress() string {
	return e.address.place
}

func main() {

	emp := &Employee{"Bala", "Krishna", Address{"chennai"}}
	fmt.Println("Employee Object value : ", emp)

	//Both statement is correct.
	// But it is not necessary to call method using pointer dereferencing variable
	//Because go will automatically convert variable in to pointer dereferencing syntax

	//(*emp).changeFirstName("balaji")
	emp.changeFirstName("balaji")
	fmt.Println("My Full Name After change is : ", emp.getFullName())

	emp.changeAddress("Bengaluru")
	fmt.Println("Address After changed  : ", emp.getAddress())

}
