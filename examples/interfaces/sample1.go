package main

import "fmt"

type Vechicle interface {
	displayWheelCount()
	displayVechileType()
}

type Car struct {
	wheelCount int
}

type auto struct {
	wheelCount int
}

func (car Car) displayWheelCount() {
	fmt.Println("Total Wheels are : ", car.wheelCount)
}

func (car Car) displayVechileType() {
	fmt.Println("Vechile Type : Car")
}

func (a auto) displayWheelCount() {
	fmt.Println("Total Wheels are : ", a.wheelCount)
}

func (a auto) displayVechileType() {
	fmt.Println("Vechile Type : Auto")
}

func main() {

	var v Vechicle
	v = Car{4}
	v.displayWheelCount()
	v.displayVechileType()

	v = auto{3}
	v.displayWheelCount()
	v.displayVechileType()
}
