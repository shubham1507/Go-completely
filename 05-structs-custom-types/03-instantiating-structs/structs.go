package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthday  string
	createdAt time.Time
}

func main() {
	userFirstName := getUserData("Please enter your first name:")
	userLastName := getUserData("Please enter your last name:")
	userBirthday := getUserData("Please enter your birthday:(MM/DD/YYYY)")

	var appUser user

	appUser = user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthday:  userBirthday,
		createdAt: time.Now(),
	}

	outputUserDetail(appUser.lastName, appUser.firstName, appUser.birthday, appUser.createdAt)
}

func outputUserDetail(lastName, firstName, birthday string, createdAt time.Time) {

	fmt.Println(firstName, lastName, birthday, createdAt)
}

func getUserData(prompt string) string {
	var input string
	fmt.Println(prompt)
	fmt.Scan(&input)
	return input
}
