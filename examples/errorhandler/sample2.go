package main

import (
	"errors"
	"fmt"
)

func divide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New(" b value should not be 0")
	}

	return a / b, nil

}

func main() {

	var i interface{}
	res, err := divide(10, 0)
	if err != nil {
		fmt.Println("Error is :", err)
	} else {
		fmt.Println(res)
	}

	res, err = divide(10, 5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Result is : ", res)
	}

}
