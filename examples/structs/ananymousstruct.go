package main

import "fmt"

func main() {

	emp := struct {
		firstName, lastName string
		salary              int
		fullTime            bool
	}{firstName: "sathya", salary: 2000}

	fmt.Println("First name : ", emp.firstName)
	fmt.Println("Last Name: ", emp.lastName)
	fmt.Println("Full Time Employee? ", emp.fullTime)
	fmt.Println("salary:", emp.salary)

	fmt.Printf("%T", emp)
}
