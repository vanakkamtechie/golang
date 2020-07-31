package main

import "fmt"

type employee struct {
	firstName, lastName string
	salary              int
	fullTime            bool
}

func main() {
	var emp employee
	emp.firstName = "bala"
	emp.lastName = "krishnan"
	emp.fullTime = true

	fmt.Println(emp.firstName)
	fmt.Println(emp.lastName)
	fmt.Println(emp.fullTime)

	fmt.Printf("%T", emp)
}
