package main

import "fmt"

func main() {
	age := 32
	var agePointer *int
	agePointer = &age
	fmt.Println("Age before edit:", agePointer)
	editAgeToOlderYears(agePointer)
	fmt.Println("Age after edit:", agePointer)
}

func editAgeToOlderYears(agePtr *int) {
	*agePtr = *agePtr - 18
}
