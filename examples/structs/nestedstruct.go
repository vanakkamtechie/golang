package main

import "fmt"

type Salary struct {
	basic     int
	allowance int
}

type employee struct {
	firstName, lastName string
	Salary              //promoted fields. Now employee can access fields defined inside Salary
	fullTime            bool
}

func main() {

	emp := &employee{"bala", "krishnan", Salary{1000, 2000}, true}

	fmt.Println("First name : ", (*emp).firstName)
	fmt.Println("Last Name: ", emp.lastName)
	fmt.Println("Full Time Employee? ", emp.fullTime)
	fmt.Println("salary:", emp.Salary)
	fmt.Println(emp)
	fmt.Println()

	//Another way of assigning values
	emp = &employee{}
	emp.firstName = "poorana"
	emp.lastName = "sathya"

	sal := Salary{}
	sal.allowance = 1000
	sal.basic = 2000
	emp.Salary = sal

	fmt.Println("First name : ", (*emp).firstName)
	fmt.Println("Last Name: ", emp.lastName)
	fmt.Println("Full Time Employee? ", emp.fullTime)
	fmt.Println("salary:", emp.Salary)
	fmt.Println("Basic:", emp.basic)         // Parent struct have variable access of chilc struct
	fmt.Println("Allowance:", emp.allowance) // Parent struct have variable access of chilc struct
	fmt.Println(emp)

}
