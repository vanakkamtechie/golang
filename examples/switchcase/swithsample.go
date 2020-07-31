package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("df")
	i := 10
	switch i {
	case 10:
		fmt.Println(">5")
		break
	default:
		fmt.Println("default")

	}

	fmt.Println("day:", time.Day()+1)

	switch time.Now().Weekday() {
	case time.Tuesday:
		fmt.Println("Tuesday")

	}

}
