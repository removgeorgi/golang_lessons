package main

import (
	"errors"
	"fmt"
)

//function that takes divident and divisor and prints the quotient

// Declaring the function, which takes two parameters and returns two(plus error!)
func divide(divident float64, divisor float64) (float64, error) {
	//Error handling - if the divisor is equal to 0 then an error message is shown. Two parameters are returned!
	if divisor == 0 {
		return 0, errors.New("can't divide by 0")
	}
	//If not the function returns the result and the default err value
	return divident / divisor, nil
}

func main() {
	//Two variables are declared to call the divide() function (Plus err!)
	quotient, err := divide(5.6, 0.0)
	//Here if the divisor is == 0 then the error is hit and printed
	if err != nil {
		fmt.Println(err)
		//If not then the formatted quotient is printed
	} else {
		fmt.Printf("%0.2f\n", quotient)
	}
}
