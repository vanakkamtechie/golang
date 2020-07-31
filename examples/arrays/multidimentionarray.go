package main

import "fmt"

func main() {
	arr := [2][2]int{{1, 2}}
	arr[0][0] = -1
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	for _, value := range arr {
		fmt.Println(" value:", value)
	}
}
