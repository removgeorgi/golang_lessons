package main

import "fmt"

func functionA(a int, b int) {
	fmt.Println(a + b)
}

func functionB(a int, b int) {
	fmt.Println(a * b)
}

func functionC(a bool) {
	fmt.Println(!a)
}

func functionD(a string, b int) {
	for i := 0; i < b; i++ {
		fmt.Print(a)
	}
}

func main() {

	functionA(2, 3)
	functionB(2, 3)
	functionC(true)
	functionD("ha", 3)

}
