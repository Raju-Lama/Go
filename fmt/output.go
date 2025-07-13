package main

import (
	"fmt"
)

func main() {
	i := 42
	j := 3.14
	fmt.Print(i, j)
	// prints without space or added new line

	fmt.Println(i, j)
	// prints with space and adds a new line

	fmt.Printf("%d %f\n", i, j)
	// formatted output with specified format

	fmt.Printf("%v, %T", i, i)
	// prints value and type of i
}
