package main

import "fmt"

var v1 = 10

func main() {

	showValue()
	fmt.Println(v1)
}

func showValue() {
	v1 := 20
	fmt.Println(v1)
}
