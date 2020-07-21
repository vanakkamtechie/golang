package main

import (
	"fmt"

	ps1 "github.com/vanakkamtechie/golang/examples/pointers"
)

// Add adds two integers
func Add(x, y int) int {
	return x + y
}

func main() {
	var i, j = 42, 50

	fmt.Println("Hello World! Result: ", Add(i, j))

	fmt.Println(&i, &j)

	ps1.DisplayFileName()
}
