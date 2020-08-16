package main

import "fmt"

func main() {
	add := func(a int, b int) int {
		return a + b
	}(10, 20)

	subtract := func(a int, b int) int {
		return a - b
	}

	fmt.Println("add2:", add)
	fmt.Println("subtract:", subtract(30, 20))

}
