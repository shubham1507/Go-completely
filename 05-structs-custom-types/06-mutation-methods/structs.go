package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (u *user) outputUserDetails() {
	fmt.Printf("User Details:\n")
	fmt.Printf("First Name: %s\n", u.firstName)
	fmt.Printf("Last Name: %s\n", u.lastName)
	fmt.Printf("Birth Date: %s\n", u.birthDate)
	fmt.Printf("Created At: %s\n", u.createdAt.Format(time.RFC1123))
}

func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}
func main() {
	userFirstName := getInput("Enter your first name: ")
	userLastName := getInput("Enter your last name: ")
	userBirthDate := getInput("Enter your birth date (YYYY-MM-DD): ")

	var newUser user

	newUser = user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthDate,
		createdAt: time.Now(),
	}
	newUser.outputUserDetails()
	newUser.clearUserName()
	newUser.outputUserDetails()
}

func getInput(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scanln(&input)
	return input
}
