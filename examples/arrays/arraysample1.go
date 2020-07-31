package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3}
	arr[2] = -1
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for _, value := range arr {
		fmt.Println(" value:", value)
	}
}
