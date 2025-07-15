package main

import (
	"fmt"
)

func main() {
	// like dictionary in python mapping is used to map key with value
	// syntax
	// var x = map[keytype]valuetype{key:value, key:value....}

	var dict = map[int]string{1: "one", 2: "two"}

	fmt.Println(dict)
}
