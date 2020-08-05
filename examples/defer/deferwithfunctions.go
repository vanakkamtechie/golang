package main

import "fmt"

func main() {
	call("method")
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
