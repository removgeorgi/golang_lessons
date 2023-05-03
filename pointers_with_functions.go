package main

import "fmt"

/*
	1.It is possible to return pointers from functions - createPointer();
	Just declare that the function's return type is a
	pointer type:
*/

// Declare that the function returns a float64 pointer:
func createPointer() *float64 {
	var myFloat = 98.5
	//Return a pointer of the specified type:
	return &myFloat
}

/*2.It is possible alsso to pass pointers to functions as arguments
Just specify that the type of one or more parameters should be a
pointer - printPointer()
*/

func printPointer(myBooleanPointer *bool) {
	fmt.Println(*myBooleanPointer)
}

func main() {

	//Assign the returned pointer to a variable:
	var myFloatPointer *float64 = createPointer()
	//Print the value at the pointer:
	fmt.Println(*myFloatPointer)

	var myBool bool = true
	printPointer(&myBool)

}
