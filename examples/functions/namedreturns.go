package main

import "fmt"

func main() {

	fmt.Println(getName())

}

func getName() (firstName, lastName, fullName string) {

	firstName = "bala"
	fullName = "Balakrishnan"
	lastName = "krishnan"

	return
}
