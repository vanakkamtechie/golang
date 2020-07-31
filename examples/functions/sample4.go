package main

import "fmt"

type Address struct {
	place string
}

type Employee struct {
	firstName, lastName string
	Address
}

func (e Employee) getFullName() string {
	return e.firstName + " " + e.lastName
}

func (e *Employee) changeFirstName(firstNameParam string) {
	e.firstName = firstNameParam
}

func (e *Address) changeAddress(place string) {
	e.place = place
}

func (e Address) getAddress() string {
	return e.place
}

func main() {

	emp := &Employee{"Bala", "Krishna", Address{"chennai"}}
	fmt.Println("Employee Object value : ", emp)

	address := Address{"Hydrabad"}
	fmt.Println("My Address : ", address.getAddress())
	fmt.Println("Before ddress After changed  : ", emp.getAddress())
	emp.Address = address
	fmt.Println("Afrer Address After changed  : ", emp.getAddress())

	emp.changeAddress("Bengaluru")
	fmt.Println("Address After changed to bengaluru  : ", emp.getAddress())

}
