package main

import (
	"fmt"
	"golang_lessons/go/src/datafile"
	"log"
)

func main() {

	numbers, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}

	fmt.Printf("Total for the last weeks is: %0.2f\n", sum)

	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)

}
