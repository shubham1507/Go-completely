package main

import (
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	Age       int
	Email     string
	CreatedAt time.Time
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	age := getUserData("Please enter your age: ")
	email := getUserData("Please enter your email: ")

	outputUserDetails(firstName, lastName, age, email)
}

func outputUserDetails(firstName, lastName, ageStr, email string) {
	fmt.Println(firstName, lastName, ageStr, email)
}

func getUserData(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scanln(&input)
	return input
}
