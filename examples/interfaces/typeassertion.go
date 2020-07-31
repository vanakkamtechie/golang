package main

import "fmt"

type Shape interface {
	Area() float64
}

type Object interface {
	volume() float64
}

type Skin interface {
	color() string
}

type Cube struct {
	side float64
}

func (c Cube) Area() float64 {
	return 6 * (c.side * c.side)
}

func (c Cube) volume() float64 {
	return c.side * c.side * c.side
}

func explain(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("\n Param is string")
	case Cube:
		fmt.Println("\n cube")
	}
}

func main() {

	//Type assertion syntex ==> value, ok := i.(Type)
	var s Shape = Cube{4}

	value, ok := s.(Object)
	fmt.Printf(" \n Type of 's' is %T and is implements interfact Object ? %v ", value, ok)

	value1, ok1 := s.(Skin)
	fmt.Printf("\nType of 's' is %T and is implements interfact Skin? %v ", value1, ok1)

	explain("vanakkam techie")
	explain(s)
}
