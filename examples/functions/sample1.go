package main

import "fmt"

type Employee struct {
	firstName, lastName string
}

func main() {
	//In functions, we should send the same type defined in the functions
	//But in functional receiver, it is not the case. Go compiler will automatically convert it
	emp := &Employee{"Bala", "Krishna"}
	displayName(emp)
}

func displayName(e *Employee) {
	fmt.Println(e)
}
