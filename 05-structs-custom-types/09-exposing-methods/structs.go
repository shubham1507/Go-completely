package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	userFirstName := getInput("Enter your first name: ")
	userLastName := getInput("Enter your last name: ")
	userBirthDate := getInput("Enter your birth date (YYYY-MM-DD): ")
	userEmail := getInput("Enter your email address: ")
	var appUser *user.User
	appUser, err := user.NewUser(userFirstName, userLastName, userBirthDate, userEmail)
	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}
	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getInput(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scanln(&input)
	return input
}
