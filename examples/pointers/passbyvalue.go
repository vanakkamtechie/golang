package main

import "fmt"

//test
func ConcatinateWordByPassingReference(s string) {
	s = *s + " World"
}

func ConcatinateWordWithPassByValue(s string) {
	s = s + " World"
}

func main() {
	var s string = "Hello"
	ConcatinateWordWithPassByValue(s)
	fmt.Println("Concatinated String: %#v", s)

	ConcatinateWordByPassingReference(&s)
	fmt.Println("After passing Reference: %v", s)

}
