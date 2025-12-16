package main

import (
	"errors"
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

func newUser(fName, lName, birthDate, email string) (*user, error) {
	if fName == "" || lName == "" {
		return nil, errors.New("first name and last name cannot be empty")
	}
	if len(email) < 5 || !containsAtSymbol(email) {
		return nil, errors.New("invalid email address")
	}
	return &user{
		fName:     fName,
		lName:     lName,
		birthDate: birthDate,
		Email:     email,
		CreatedAt: time.Now(),
	}, nil
}

// func containsAtSymbol(email string) bool {
// 	panic("unimplemented")
// }

func main() {
	userFirstName := getInput("Enter your first name: ")
	userLastName := getInput("Enter your last name: ")
	userBirthDate := getInput("Enter your birth date (YYYY-MM-DD): ")
	userEmail := getInput("Enter your email address: ")
	var appUser *user
	var err error
	appUser, err = newUser(userFirstName, userLastName, userBirthDate, userEmail)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}
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
func containsAtSymbol(email string) bool {
	for _, char := range email {
		if char == '@' {
			return true
		}
	}
	return false
}
