package main

import "fmt"

func main() {
	fmt.Println(addTwoValues(20, 10))
	add, sub, div := getAddMulDivideValue(20, 10)
	fmt.Printf("add: %d, sub:%d, div:%f", add, sub, div)
}

func addTwoValues(a int, b int) int {
	return a + b
}

func getAddMulDivideValue(val1 int, val2 int) (int, int, float32) {
	add := val1 + val2
	mul := val1 * val2
	div := val1 / val2

	return add, mul, float32(div)

}
