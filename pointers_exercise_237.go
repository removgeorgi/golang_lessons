package main

import (
	"fmt"
	"reflect"
)

func createPointer() *float64 {
	var myFloat = 98.5
	return &myFloat

}

func main() {

	var myFloatPointer *float64 = createPointer()
	fmt.Println(*myFloatPointer)
	fmt.Println(reflect.TypeOf(*myFloatPointer))

	// Declaring myInt int value
	var myInt int
	// Declaring myIntPointer as a value that holds an integer pointer
	myIntPointer := &myInt
	//Assigning integer pointer value to myIntPointer
	*myIntPointer = 42
	//Prints the integer pointer value
	fmt.Println(*myIntPointer)
}
