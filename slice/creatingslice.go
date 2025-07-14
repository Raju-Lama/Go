package main

import (
	"fmt"
)

func main() {
	slice1 := []int{}
	//empty slice
	fmt.Println(cap(slice1)) // 0
	// for capacity
	fmt.Println(len(slice1)) // 0
	//actual length

	slice2 := []int{1, 2, 3, 4, 5}
	//length 5, capacity 5
	fmt.Println(cap(slice2)) // 5
	fmt.Println(len(slice2)) // 5

	arr := [3]int{1, 2, 3}

	slice3 := arr[0:2]
	//slice of array, length 2, capacity 3
	fmt.Println(cap(slice3)) // 3
	fmt.Println(len(slice3)) // 2
}
