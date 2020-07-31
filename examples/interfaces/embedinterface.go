package main

import "fmt"

type Shape interface {
	Area() float64
}

type Object interface {
	Volume() float64
}

//Interface cannot implement other interface. But we can embed it together as we are doing it for struct
type Material interface {
	Shape
	Object
}

type Cube struct {
	side float64
}

func (c Cube) Area() float64 {
	return 6 * (c.side * c.side)
}

func (c Cube) Volume() float64 {
	return c.side * c.side * c.side
}

func main() {
	c := Cube{3}
	var m Material = c
	var s Shape = c
	var o Object = c
	fmt.Printf("dynamic type and value of interface m of static type Material is'%T' and '%v'\n", m, m)
	fmt.Printf("dynamic type and value of interface s of static type Shape is'%T' and '%v'\n", s, s)
	fmt.Printf("dynamic type and value of interface o of static type Object is'%T' and '%v'\n", o, o)
}
