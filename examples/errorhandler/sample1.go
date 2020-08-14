package main

import (
	"errors"
	"fmt"
)

type MyError struct{}

func (e *MyError) Error() string {
	return "Unexpected Error happened"
}

func main() {
	var e = &MyError{}
	fmt.Println(e)
	var err = errors.New("New Buildin type")
	fmt.Println(err)
}
