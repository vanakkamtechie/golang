package main

import "fmt"

func main() {
	i := 0
do:
	fmt.Println("dowhile: i value is ", i)

	for i < 10 {
		i++
		goto do
	}
}
