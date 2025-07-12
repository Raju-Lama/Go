package main

import (
	"fmt"
)

func main() {
	a := 10
	// Implicit type conversion in Go
	b := 20.5
	c := "Hello, Go!"
	fmt.Println("Integer:", a)
	fmt.Println("Float:", b)
	fmt.Println("String:", c)
	fmt.Printf("Type of a: %T\n", a)
	fmt.Printf("Type of b: %T\n", b)
	fmt.Printf("Type of c: %T\n", c)
	fmt.Println("Sum of a and b:", a+int(b)) // Explicit conversion of float

	var d float64 = 30.5
	fmt.Println("Sum of b and d:", b+d)                        // Implicit conversion of int to float
	fmt.Println("Concatenation of c and string:", c+" World!") // Implicit conversion
}
