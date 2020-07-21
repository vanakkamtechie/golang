package main

import (
	"fmt"
	"reflect"
)

func main() {

	//DataType
	anInt := 1234
	fmt.Printf(" Datatype of %d is %T\n", anInt, anInt)
	fmt.Printf(reflect.TypeOf(anInt).String())
}
