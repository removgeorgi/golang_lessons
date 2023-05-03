package main

import "fmt"

func main() {
	var myInt int
	// Declare a variable that holds a pointer to an int:
	var myIntPointer *int
	// Assign a pointer to the variable:
	myIntPointer = &myInt
	fmt.Println(myIntPointer)

	// The same with float64 variable:
	var myFloat float64
	var myFloatPointer *float64
	myFloatPointer = &myFloat
	fmt.Println(myFloatPointer)

	// A short declaration for a pointer variable:
	var myBool bool
	myBoolPointer := &myBool
	fmt.Println(myBoolPointer)

	//Getting or changing the value at a pointer

	myInt2 := 4
	myInt2Pointer := &myInt2
	// Print the pointer itself:
	fmt.Println(myInt2Pointer)
	// Print the value at the pointer
	fmt.Println(*myInt2Pointer)

	myFloat2 := 98.6
	myFloat2Pointer := &myFloat2
	fmt.Println(myFloat2Pointer)
	fmt.Println(*myFloat2Pointer)

	myBool2 := true
	myBool2Pointer := &myBool2
	fmt.Println(myBool2Pointer)
	fmt.Println(*myBool2Pointer)

	/*
		****************************
		The * operator can also be used to update a value at a pointer
		****************************
	*/

	myInt3 := 4
	fmt.Println(myInt3)
	myInt3Pointer := &myInt3
	//Assign a new value to the variable at the pointer (myInt3):
	*myInt3Pointer = 8
	//Print the value of the variable at the pointer:
	fmt.Println(*myInt3Pointer)
	//Print the variable's value directly:
	fmt.Println(myInt3)
}
