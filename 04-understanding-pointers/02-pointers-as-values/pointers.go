package main

import "fmt"

func main() {

	age := 25

	var agePointer *int

	agePointer = &age

	fmt.Println("Age", *agePointer)
}
