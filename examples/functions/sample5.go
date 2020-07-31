package main

import (
	"fmt"
	"strings"
)

type MyString string

/*func (s string) toUpperCase() string {
	return strings.ToUpper(s)
}*/

func (s MyString) toUpperCase() string {
	tmpString := string(s)
	return strings.ToUpper(tmpString)
}

func main() {
	//s := "Hello World"
	s := MyString("Hello World")
	fmt.Println(s.toUpperCase())
}
