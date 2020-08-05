package main

import "fmt"

func main() {
	s := adder()

	fmt.Println(s(10))
	fmt.Println(s(20))
}

func adder() func(int) int {
	var sum int
	return func(val int) int {
		sum = sum + val
		return sum
	}
}
