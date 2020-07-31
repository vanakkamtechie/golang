package main

import "fmt"

type MyFullNameFunction func(string, string) string

type employee struct {
	firstName string
	lastName  string
	fullName  MyFullNameFunction
}

func main() {

	e := employee{"bala", "krishnan", func(fn string, ln string) string { return fn + " " + ln }}
	fmt.Println("My fullName: ", e.fullName(e.firstName, e.lastName))
}
