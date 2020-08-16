package main

import "fmt"

func main() {
	fmt.Println("add:", calculate(20, 10, add))
	fmt.Println("subtract:", calculate(20, 10, subtract))

}

func add(a, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func calculate(a int, b int, f func(int, int) int) int {
	val := f(a, b)
	return val
}
