package main

import "fmt"

func main() {

	age := 30
	fmt.Println("Age:", age)

	adultAge := getAdultAge(age)
	fmt.Println("Adult Age:", adultAge)
}

func getAdultAge(age int) int {

	return age - 18
}
