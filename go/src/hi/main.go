package main

import (

	//"golang_lessons/go/src/keyboard"
	"golang_lessons/go/src/greeting"
	"golang_lessons/go/src/greeting/deutsch"
)

func main() {

	//The commented code below is related to run function from the keyboard package:
	// 	fmt.Print("Enter a grade:")
	// 	grade, err := keyboard.GetFloat()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	var status string

	// 	if grade >= 60 {
	// 		status = "passing"
	// 	} else {
	// 		status = "failing"
	// 	}

	// 	fmt.Println("A grade of", grade, "is", status)

	greeting.Hello()
	greeting.Hi()
	deutsch.Hallo()
	deutsch.GutenTag()

}
