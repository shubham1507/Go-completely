package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	fName     string
	lName     string
	birthDate string
	Email     string
	CreatedAt time.Time
}

func (u *User) OutputUserDetails() {
	fmt.Println(u.fName, u.lName, u.birthDate)
}
func (u *User) ClearUserName() {
	u.fName = ""
	u.lName = ""
}

func NewUser(fName, lName, birthDate, email string) (*User, error) {
	if fName == "" || lName == "" {
		return nil, errors.New("first name and last name cannot be empty")
	}
	if len(email) < 5 || !containsAtSymbol(email) {
		return nil, errors.New("invalid email address")
	}
	return &User{
		fName:     fName,
		lName:     lName,
		birthDate: birthDate,
		Email:     email,
		CreatedAt: time.Now(),
	}, nil
}
func containsAtSymbol(email string) bool {
	for _, char := range email {
		if char == '@' {
			return true
		}
	}
	return false
}
