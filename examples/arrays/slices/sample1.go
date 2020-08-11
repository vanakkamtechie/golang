package main

import "fmt"

func main() {
	var arr = [5]int{10, 20, 30, 40, 50}
	//displayValue(arr...)
	for k, v := range arr {
		fmt.Println("array values: key:", k, " val:", v)
	}

	fmt.Println()

	var s = arr[2:3]
	displayValue(s...)

	arr[2] = 100

	fmt.Println(&arr[2])
	fmt.Println(&s[0])

	s = append(s, 60, 70, 80, 90, 100)

	fmt.Println(arr)
	fmt.Println(s)
	s[0] = 110

	fmt.Println(&arr[0])
	fmt.Println(&s[0])
}

func displayValue(arr ...int) {
	for k, v := range arr {
		fmt.Println("slice values: key:", k, " val:", v)
	}
}
