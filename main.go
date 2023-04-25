package main

import (
	"fmt"
	"strings"
)

func main() {

	/*now := time.Now()
	year := now.Year()
	fmt.Println(now)
	fmt.Println(year)
	*/

	broken := "M& birthda& is coming"
	replacer := strings.NewReplacer("&", "y")
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)

}
