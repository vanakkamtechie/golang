package main

import "fmt"

func main() {
	sum(10, 12, 34)
}

func sum(numbers ...int) {
	for _, v := range numbers {
		fmt.Println(v)
	}
}
