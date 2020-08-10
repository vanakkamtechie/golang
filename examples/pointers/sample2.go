package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	var v *Vertex
	fmt.Printf("Vertex %p \n", v)

	var p = new(Vertex)
	p.X = 10
	fmt.Println(p.X)

	p1 := p
	p1.X = 20

	fmt.Println(p.X)

	changeXValue(p1)

	fmt.Println(p.X)

}

func changeXValue(vert *Vertex) {
	vert.X = 100
}
