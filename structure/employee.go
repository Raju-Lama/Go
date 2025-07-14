package main

import (
	"fmt"
)

type Employee struct {
	id   int
	name string
	age  int
}

func main() {
	var emp Employee

	emp.id = 1
	emp.name = "John Doe"
	emp.age = 30

	fmt.Println("Employee ID:", emp.id)
	fmt.Println("Employee Name:", emp.name)
	fmt.Println("Employee Age:", emp.age)

	Person(emp)

}

func Person(em Employee) {
	fmt.Println("Employee Details:")
	fmt.Printf("ID: %d, Name: %s, Age: %d\n", em.id, em.name, em.age)
}
