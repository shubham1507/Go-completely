package main

import "fmt"

func main() {
	firstName := "Shubham"
	lastName := "Joshi"
	fullname := fmt.Sprintf("%v %v", firstName, lastName)
	age := 31
	fmt.Printf(" Hi , my name is %v and I am %v yrs old.", fullname, age)
}
