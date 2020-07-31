package main

import "fmt"

type Car struct {
	wheelCount int
}

type auto struct {
	wheelCount int
}

type MyString string

func displayValues(i interface{}) {
	fmt.Printf("Value type: %T, value: %v\n", i, i)
}

func main() {

	v := Car{4}
	displayValues(v)
	displayValues(MyString("customString"))
}
