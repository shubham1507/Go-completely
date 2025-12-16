package main

import (
	"fmt"
	"strconv"
	"time"
)

type user struct {
	firstName string
	lastName  string
	age       int
	createdAt time.Time
}

func (u user) outputUserInfo() {
	fmt.Printf("User Info:\n")
	fmt.Println(u.firstName, u.lastName, u.age)
}
func main() {

	userFirstName := getInput("Enter your first name: ")
	userLastName := getInput("Enter your last name: ")
	userAgeStr := getInput("Enter your age: ")
	userAge, _ := strconv.Atoi(userAgeStr)

	var appUser user

	appUser = user{
		firstName: userFirstName,
		lastName:  userLastName,
		age:       userAge,
		createdAt: time.Now(),
	}

	appUser.outputUserInfo()
}

func getInput(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scanln(&input)
	return input
}
