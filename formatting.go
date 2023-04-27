package main

import (
	"fmt"
	"log"
)

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}
	area := width * height
	return area / 10.0, nil

}

func main() {

	var amount, total float64

	amount, err := paintNeeded(4.2, 3.0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%0.2f liters needed for the first wall\n", amount)
	total += amount

	amount, err = paintNeeded(5.2, -3.5)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%0.2f liters needed for the second wall\n", amount)
	total += amount

	fmt.Printf("Total: %0.2f liters\n", total)

	// fmt.Printf("A float: %f\n", 3.1415)
	// fmt.Printf("An integer: %d\n", 15)
	// fmt.Printf("A string: %s\n", "hello")
	// fmt.Printf("A boolean: %t\n", false)
	// fmt.Printf("Values: %v %v %v\n", 1.2, "\t", true)
	// fmt.Printf("Values: %#v %#v %#v\n", 1.2, "\t", true)
	// fmt.Printf("Types: %T %T %T\n", 1.2, "\t", true)
	// fmt.Printf("Percent sign: %%\n")

	// fmt.Printf("%12s | %s\n", "Product", "Cost in Cents")
	// fmt.Println("---------------------------------------")

	// fmt.Printf("%12s | %2d\n", "Stamps", 50)
	// fmt.Printf("%12s | %2d\n", "Paper Clips", 5)
	// fmt.Printf("%12s | %2d\n", "Tape", 99)

	// paintNeeded(4.2, 3.0)
	// paintNeeded(5.2, 3.5)
	// paintNeeded(5.0, 3.3)

	// fmt.Println(status(59))
	// fmt.Println(status(100))

}
