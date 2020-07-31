package main

import "fmt"

type employee struct {
	firstName, lastName string
	salary              int
	fullTime            bool
}

func main() {

	emp := &employee{"bala", "krishnan", 2000, true}

	fmt.Println("First name : ", (*emp).firstName)
	fmt.Println("Last Name: ", emp.lastName)
	fmt.Println("Full Time Employee? ", emp.fullTime)
	fmt.Println("salary:", emp.salary)
	fmt.Println(emp)

}
