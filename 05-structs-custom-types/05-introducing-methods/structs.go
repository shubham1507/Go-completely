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

func (u user) outputUserDetails() {

	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func main() {
	userFirstName := getUserData(" Please enter your fname: \n")
	userLastName := getUserData(" Please enter your lname: \n")
	userBirthdate := getUserData(" Please enter your birthdate (DD/MM/YYYY): \n")

	appUser := user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthdate,
		createdAt: time.Now(),
	}

	appUser.outputUserDetails()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
