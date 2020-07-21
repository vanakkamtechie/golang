package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	var v *Vertex
	fmt.Printf("Vertex %p \n", v)

	var p = *new(Vertex)
	fmt.Printf("Vertex %p \n", p)
}
