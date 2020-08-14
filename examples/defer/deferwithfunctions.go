package main

import "fmt"

func main() {
	defer fmt.Println("progrm execution completed")
	defer call("end")
	fmt.Println("Main function completed")
}

func call(param string) {
	defer func() {
		fmt.Println("Calling defer method...")
	}()
	fmt.Println("Bye..")
	if param == "end" {
		fmt.Println(param)
	}
}
