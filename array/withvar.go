package main

import (
	"fmt"
)

func main() {
	// declare with var keyword
	var ar = [5]int{1, 2, 3, 4, 5}
	fmt.Println(ar)

	// declare with shorthand syntax
	ar2 := [...]int{6, 7, 8, 9, 10}
	fmt.Println(ar2)

}
