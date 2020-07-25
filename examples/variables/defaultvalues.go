package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	//https://golang.org/pkg/fmt/

	//DataType
	anInt := 1234
	fmt.Printf(" Datatype of %d is %T\n", anInt, anInt)

	//anInt = "stringvalue"
	//fmt.Printf(" Datatype of %d is %T\n", anInt, anInt)

	// Will print 'string'
	aString := "Hello World"
	fmt.Printf(" Datatype of %s is %T\n", aString, aString)

	// Will print 'float64'
	aFloat := 3.14
	fmt.Printf(" Datatype of %f is %T\n\n\n", aFloat, aFloat)

	//Default Value
	var i int
	fmt.Printf(" Default Value of Interger: %#v \n", i)

	var s string
	fmt.Printf(" Default Value of String :  %#v \n", s)

	var ps *string
	fmt.Printf(" Default Value of pointer String :  %#v \n", ps)

	var b bool
	fmt.Printf(" Default Value of bool : %#v \n", b)

	var t Vertex
	fmt.Printf(" Default Value of struct : %#v \n", t)

	var p *Vertex
	fmt.Printf(" Default Value of struct : %#v \n", p)

	if p != nil {
		fmt.Println(" Vertex is not null")
	} else {
		fmt.Println(" Vertex is null")
	}

	/* nil
	-- pointer
	   function
	   interface
	   slices
	   channels
	   maps 
	   */

}
