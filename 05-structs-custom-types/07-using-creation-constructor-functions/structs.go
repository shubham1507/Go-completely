package main

import (
	"fmt"
	"time"
)

type user struct {
	fName     string
	lName     string
	birthDate string
	Email     string
	CreatedAt time.Time
}

func (u *user) outputUserDetails() {
	fmt.Println(u.fName, u.lName, u.birthDate)
}

func (u *user) clearUserName() {
	u.fName = ""
	u.lName = ""
}

func newUser(fName, lName, birthDate, email string) *user {
	return &user{
		fName:     fName,
		lName:     lName,
		birthDate: birthDate,
		Email:     email,
		CreatedAt: time.Now(),
	}
}

func main() {
	userFirstName := getInput("Enter your first name: ")
	userLastName := getInput("Enter your last name: ")
	userBirthDate := getInput("Enter your birth date (YYYY-MM-DD): ")
	userEmail := getInput("Enter your email address: ")

	var appUser *user
	appUser = newUser(userFirstName, userLastName, userBirthDate, userEmail)
	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()
}

func getInput(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scanln(&input)
	return input
}
